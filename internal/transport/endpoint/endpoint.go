package endpoint

import (
	"biterush"
	"biterush/generated/database/model"
	"context"

	"k8s.io/klog/v2"
)

type BiteEndpoint struct {
	BiteManager biterush.Bite
}

type OrderRequest struct {
	UserID       *int64
	RiderID      *int64
	RestaurantID *int64
	Status       *biterush.Status
}

type RiderRequest struct {
	RiderID   int64
	Latitude  *float64
	Longitude *float64
}

type UpdateOrderRequest struct {
	OrderID int64
	Status  string
}

func (ep *BiteEndpoint) CreateUser(ctx context.Context, request interface{}) (interface{}, error) {
	params := request.(*model.Users)

	logger := klog.LoggerWithName(klog.FromContext(ctx), "Create User")
	ctx = klog.NewContext(ctx, logger)

	return ep.BiteManager.CreateUser(ctx, params)
}

func (ep *BiteEndpoint) CreateRider(ctx context.Context, request interface{}) (interface{}, error) {
	params := request.(*model.Riders)

	logger := klog.LoggerWithName(klog.FromContext(ctx), "Create Rider")
	ctx = klog.NewContext(ctx, logger)

	return ep.BiteManager.CreateRider(ctx, params)
}

func (ep *BiteEndpoint) CreateRestaurant(ctx context.Context, request interface{}) (interface{}, error) {
	params := request.(*model.Restaurants)

	logger := klog.LoggerWithName(klog.FromContext(ctx), "Create Restaurant")
	ctx = klog.NewContext(ctx, logger)

	return ep.BiteManager.CreateRestaurant(ctx, params)
}

func (ep *BiteEndpoint) CreateOrder(ctx context.Context, request interface{}) (interface{}, error) {
	params := request.(*model.Orders)

	logger := klog.LoggerWithName(klog.FromContext(ctx), "Create Order")
	ctx = klog.NewContext(ctx, logger)

	return ep.BiteManager.CreateOrder(ctx, params)
}

func (ep *BiteEndpoint) GetMenu(ctx context.Context, request interface{}) (interface{}, error) {
	params := request.(int64)

	logger := klog.LoggerWithName(klog.FromContext(ctx), "Get Menu")
	ctx = klog.NewContext(ctx, logger)

	return ep.BiteManager.GetMenu(ctx, params)
}

func (ep *BiteEndpoint) ListOrders(ctx context.Context, request interface{}) (interface{}, error) {
	params := request.(OrderRequest)

	logger := klog.LoggerWithName(klog.FromContext(ctx), "Get Menu")
	ctx = klog.NewContext(ctx, logger)

	return ep.BiteManager.ListOrders(ctx, params.UserID, params.RiderID, params.RestaurantID, params.Status)
}

func (ep *BiteEndpoint) UpdateOrder(ctx context.Context, request interface{}) (interface{}, error) {
	params := request.(UpdateOrderRequest)

	logger := klog.LoggerWithName(klog.FromContext(ctx), "Accept Order")
	ctx = klog.NewContext(ctx, logger)

	err := ep.BiteManager.UpdateOrder(ctx, params.OrderID, params.Status)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (ep *BiteEndpoint) UpdateRider(ctx context.Context, request interface{}) (interface{}, error) {
	params := request.(RiderRequest)

	logger := klog.LoggerWithName(klog.FromContext(ctx), "Update Rider")
	ctx = klog.NewContext(ctx, logger)

	return ep.BiteManager.UpdateRider(ctx, params.RiderID, params.Latitude, params.Longitude)
}

func (ep *BiteEndpoint) FindRestaurant(ctx context.Context, request interface{}) (interface{}, error) {
	params := request.(*model.Restaurants)

	logger := klog.LoggerWithName(klog.FromContext(ctx), "Find restaurant")
	ctx = klog.NewContext(ctx, logger)

	return ep.BiteManager.FindRestaurant(ctx, params)
}
