package main

import (
	"biterush/cmd"
	"biterush/internal"
	"log"

	"biterush/generated/database/table"

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
}
