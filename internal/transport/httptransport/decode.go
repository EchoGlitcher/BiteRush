package httptransport

import (
	"biterush/generated/database/model"
	"biterush/internal/transport/endpoint"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.Users
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func decodeCreateRiderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.Riders
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func decodeCreateRestaurantRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.Restaurants
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func decodeCreateOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.Orders
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func decodeListOrdersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil && err != http.ErrBodyNotAllowed {
		// Body might be empty for GET, so it's fine
		return nil, err
	}
	return req, nil
}

func decodeAcceptOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	orderID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, err
	}
	return orderID, nil
}

func decodeUpdateRiderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.RiderRequest
	vars := mux.Vars(r)
	idStr := vars["id"]
	riderID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, err
	}
	req.RiderID = riderID
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetMenuRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	restaurantID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, err
	}
	return restaurantID, nil
}
