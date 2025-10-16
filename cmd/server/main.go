package main

import (
	"biterush"
	"biterush/bite"
	"biterush/cmd"
	"biterush/internal"
	"biterush/internal/transport/endpoint"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"biterush/generated/database/table"
	"biterush/internal/transport/httptransport"

	"github.com/gorilla/mux"
	"k8s.io/klog/v2"
)

func main() {
	config := cmd.ResolveConfig()

	db, err := internal.CreateMysqlConnection(config)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	table.UseSchema(config.Database.Name)

	defer func() {
		if err := db.Close(); err != nil {
			klog.Errorf("cannot close db connection: %v", err)
		}
	}()

	var biteManager biterush.Bite
	biteManager = bite.NewManager(db)

	biteEndpoints := &endpoint.BiteEndpoint{BiteManager: biteManager}
	httpServer := &http.Server{
		Handler: httpHandler(biteEndpoints),
	}
	{
		defer func() {
			klog.Info("stopping http server")
			if err := httpServer.Shutdown(context.Background()); err != nil {
				klog.Errorf("cannot close http server: %v", err)
			}
		}()

		go func() {
			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Server.Port))
			if err != nil {
				klog.Fatalf("cannot create http listener: %v", err)
			}

			klog.Infof("preparing to start http server on [%s]", lis.Addr().String())
			if err := httpServer.Serve(lis); err != nil && !errors.Is(err, http.ErrServerClosed) {
				klog.Fatalf("cannot start http server: %v", err)
			}
		}()
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	klog.Infof("received shutdown signal [%v]", <-sig)
}

func httpHandler(biteEndpoints *endpoint.BiteEndpoint) http.Handler {
	router := mux.NewRouter()

	biteRouter := router.PathPrefix("/biterush").Subrouter()

	httptransport.AddBiteRoutes(biteRouter, biteEndpoints)

	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}

		methods, err := route.GetMethods()
		if err != nil {
			return nil
		}

		klog.Infof("\t%v %s\n", methods, path)

		return nil
	})
	if err != nil {
		klog.Errorf("cannot print routes: %v", err)
	}

	return router
}
