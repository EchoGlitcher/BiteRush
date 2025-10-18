package httptransport

import (
	"biterush/internal/transport/endpoint"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func AddBiteRoutes(r *mux.Router, ep *endpoint.BiteEndpoint) {

	// --- USERS ---
	r.Methods("POST").Path("/users").Handler(kithttp.NewServer(
		ep.CreateUser,
		decodeCreateUserRequest,
		kithttp.EncodeJSONResponse,
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	))

	// --- RIDERS ---
	r.Methods("POST").Path("/riders").Handler(kithttp.NewServer(
		ep.CreateRider,
		decodeCreateRiderRequest,
		kithttp.EncodeJSONResponse,
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	))

	r.Methods("PUT").Path("/riders/{id}/location").Handler(kithttp.NewServer(
		ep.UpdateRider,
		decodeUpdateRiderRequest,
		kithttp.EncodeJSONResponse,
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	))

	// --- RESTAURANTS ---
	r.Methods("POST").Path("/restaurants").Handler(kithttp.NewServer(
		ep.CreateRestaurant,
		decodeCreateRestaurantRequest,
		kithttp.EncodeJSONResponse,
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	))

	r.Methods("GET").Path("/restaurants").Handler(kithttp.NewServer(
		ep.FindRestaurant,
		decodeFindRestaurantsRequest,
		kithttp.EncodeJSONResponse,
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	))

	// --- ORDERS ---
	r.Methods("POST").Path("/orders").Handler(kithttp.NewServer(
		ep.CreateOrder,
		decodeCreateOrderRequest,
		kithttp.EncodeJSONResponse,
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	))

	r.Methods("GET").Path("/orders").Handler(kithttp.NewServer(
		ep.ListOrders,
		decodeListOrdersRequest,
		kithttp.EncodeJSONResponse,
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	))

	r.Methods("PUT").Path("/orders/{id}").Handler(kithttp.NewServer(
		ep.UpdateOrder,
		decodeUpdateOrderRequest,
		kithttp.EncodeJSONResponse,
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	))

	// --- MENU ---
	r.Methods("GET").Path("/restaurants/{id}/menu").Handler(kithttp.NewServer(
		ep.GetMenu,
		decodeGetMenuRequest,
		kithttp.EncodeJSONResponse,
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	))
}
