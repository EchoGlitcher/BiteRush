package biterush

import (
	"biterush/generated/database/model"
	"context"
)

type Status string

const (
	Pending   Status = "pending"
	Accepted  Status = "accepted"
	PickedUp  Status = "PickedUp"
	Completed Status = "completed"
)

type OrderHistory struct {
	Order      model.Orders
	OrderItems []*model.OrderItems
}

type Bite interface {
	CreateUser(ctx context.Context, user *model.Users) (*model.Users, error)
	CreateRider(ctx context.Context, rider *model.Riders) (*model.Riders, error)
	CreateRestaurant(ctx context.Context, restaurant *model.Restaurants) (*model.Restaurants, error)
	CreateOrder(ctx context.Context, order *model.Orders) (*model.Orders, error)
	GetMenu(ctx context.Context, restaurantID int64) ([]*model.MenuItems, error)
	UpdateRider(ctx context.Context, riderID int64, latitude, longitude *float64) (*model.Riders, error)
	FindRestaurant(ctx context.Context, restaurant *model.Restaurants) ([]*model.Restaurants, error)
	ListOrders(ctx context.Context, userID, riderID, restaurantID *int64, status *Status) ([]*OrderHistory, error)
	UpdateOrder(ctx context.Context, orderID int64, status string) error
	FindNearestRider(ctx context.Context, restaurantID int64) ([]*model.Riders, error)
}
