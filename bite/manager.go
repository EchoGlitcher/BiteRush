package bite

import (
	"biterush"
	"biterush/generated/database/model"
	"context"
	"database/sql"
	"fmt"
	"math/rand/v2"
	"strings"
)

func NewManager(db *sql.DB) biterush.Bite {
	return &defaultManager{db: &dbOperations{db}}
}

type defaultManager struct {
	db *dbOperations
}

func (m *defaultManager) CreateUser(ctx context.Context, users *model.Users) (*model.Users, error) {
	return m.db.createUser(ctx, users)
}

func (m *defaultManager) CreateRider(ctx context.Context, rider *model.Riders) (*model.Riders, error) {
	return m.db.createRider(ctx, rider)
}

func (m *defaultManager) CreateRestaurant(ctx context.Context, restaurant *model.Restaurants) (*model.Restaurants, error) {
	return m.db.createRestaurant(ctx, restaurant)
}

func (m *defaultManager) CreateOrder(ctx context.Context, order *model.Orders) (*model.Orders, error) {
	return m.db.createOrder(ctx, order)
}

func (m *defaultManager) GetMenu(ctx context.Context, restaurantId int64) ([]*model.MenuItems, error) {
	return m.db.getMenu(ctx, restaurantId)
}

func (m *defaultManager) UpdateRider(ctx context.Context, riderID int64, latitude, longitude *float64) (*model.Riders, error) {
	return m.db.updateRider(ctx, riderID, latitude, longitude)
}

func (m *defaultManager) FindRestaurant(ctx context.Context, restaurant *model.Restaurants) ([]*model.Restaurants, error) {
	return m.db.findRestaurant(ctx, restaurant)
}

func (m *defaultManager) ListOrders(ctx context.Context, userID, riderID, restaurantID *int64, status *biterush.Status) ([]*biterush.OrderHistory, error) {
	return m.db.listOrders(ctx, userID, riderID, restaurantID, status)
}

func (m *defaultManager) UpdateOrder(ctx context.Context, orderID int64, status string) error {
	orders, err := m.db.getOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	switch strings.ToLower(status) {
	case "accepted":
		orders.Status = string(biterush.Accepted)

		riders, err := m.FindNearestRider(ctx, orders.RestaurantID)
		if err != nil {
			return err
		}

		if len(riders) == 0 {
			return fmt.Errorf("no available riders found")
		}

		selectedRider := riders[rand.Int64N(int64(len(riders)))]
		orders.RiderID = &selectedRider.ID

	case "picked_up":
		orders.Status = string(biterush.PickedUp)
	case "completed":
		orders.Status = string(biterush.Completed)
	default:
		return fmt.Errorf("invalid status: %s", status)
	}

	_, err = m.db.updateOrder(ctx, orders)
	if err != nil {
		return err
	}

	return nil
}

func (m *defaultManager) FindNearestRider(ctx context.Context, restaurantID int64) ([]*model.Riders, error) {
	restaurant, err := m.db.getRestaurantByID(ctx, restaurantID)
	if err != nil {
		return nil, err
	}
	return m.db.findNearestAvailableRider(ctx, restaurant.Latitude, restaurant.Longitude)
}
