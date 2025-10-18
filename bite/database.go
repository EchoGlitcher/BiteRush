package bite

import (
	"biterush"
	"biterush/generated/database/model"
	. "biterush/generated/database/table"
	"biterush/internal"
	"context"
	"database/sql"
	"fmt"
	"time"

	. "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
)

type dbOperations struct {
	db *sql.DB
}

type OrderHistoryModel struct {
	model.Orders
	model.OrderItems
}

func (d *dbOperations) createUser(ctx context.Context, users *model.Users) (*model.Users, error) {
	logger := klog.FromContext(ctx)

	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot start db transaction")
	}

	defer func() {
		if err != nil {
			logger.Error(err, "CreateUsers(...) transaction commit error")
			if err := tx.Rollback(); err != nil {
				logger.Error(err, "cannot rollback transaction created for CreateUsers(..)")
			}
		}
	}()

	now := time.Now()
	usersInfo := model.Users{
		Name:      users.Name,
		Email:     users.Email,
		Phone:     users.Phone,
		Address:   users.Address,
		Latitude:  users.Latitude,
		Longitude: users.Longitude,
		CreatedAt: &now,
	}

	stmt := Users.INSERT(Users.AllColumns).MODEL(usersInfo)
	internal.LogStatement(ctx, stmt)

	result, err := stmt.ExecContext(ctx, tx)
	if err != nil {
		var mErr *mysql.MySQLError
		if errors.As(err, &mErr) {
			if mErr.Number == 1062 {
				return nil, err
			}
		}

		return nil, err
	}

	userID, _ := result.LastInsertId()
	usersInfo.ID = userID

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &usersInfo, nil
}

func (d *dbOperations) createRider(ctx context.Context, rider *model.Riders) (*model.Riders, error) {
	logger := klog.FromContext(ctx)

	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot start db transaction")
	}

	defer func() {
		if err != nil {
			logger.Error(err, "CreateRiders(...) transaction commit error")
			if err := tx.Rollback(); err != nil {
				logger.Error(err, "cannot rollback transaction created for CreateRiders(..)")
			}
		}
	}()

	now := time.Now()
	ridersInfo := model.Riders{
		Name:          rider.Name,
		Email:         rider.Email,
		Phone:         rider.Phone,
		VehicleType:   rider.VehicleType,
		LicenseNumber: rider.LicenseNumber,
		Latitude:      rider.Latitude,
		Longitude:     rider.Longitude,
		IsAvailable:   rider.IsAvailable,
		CreatedAt:     &now,
	}

	stmt := Riders.INSERT(Riders.AllColumns).MODEL(ridersInfo)
	internal.LogStatement(ctx, stmt)

	result, err := stmt.ExecContext(ctx, tx)
	if err != nil {
		var mErr *mysql.MySQLError
		if errors.As(err, &mErr) {
			if mErr.Number == 1062 {
				return nil, err
			}
		}

		return nil, err
	}

	riderID, _ := result.LastInsertId()
	ridersInfo.ID = riderID

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &ridersInfo, nil
}

func (d *dbOperations) createRestaurant(ctx context.Context, restaurant *model.Restaurants) (*model.Restaurants, error) {
	logger := klog.FromContext(ctx)

	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot start db transaction")
	}

	defer func() {
		if err != nil {
			logger.Error(err, "CreateRestaurant(...) transaction commit error")
			if err := tx.Rollback(); err != nil {
				logger.Error(err, "cannot rollback transaction created for CreateRestaurant(..)")
			}
		}
	}()

	now := time.Now()
	restaurantInfo := model.Restaurants{
		Name:            restaurant.Name,
		Email:           restaurant.Email,
		Phone:           restaurant.Phone,
		Address:         restaurant.Address,
		Latitude:        restaurant.Latitude,
		Longitude:       restaurant.Longitude,
		CuisineType:     restaurant.CuisineType,
		PreparationTime: restaurant.PreparationTime,
		IsActive:        restaurant.IsActive,
		CreatedAt:       &now,
	}

	stmt := Restaurants.INSERT(Restaurants.AllColumns).MODEL(restaurantInfo)
	internal.LogStatement(ctx, stmt)

	result, err := stmt.ExecContext(ctx, tx)
	if err != nil {
		var mErr *mysql.MySQLError
		if errors.As(err, &mErr) {
			if mErr.Number == 1062 {
				return nil, err
			}
		}

		return nil, err
	}

	restaurantID, _ := result.LastInsertId()
	restaurantInfo.ID = restaurantID

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &restaurantInfo, nil
}

func (d *dbOperations) createOrder(ctx context.Context, order *model.Orders) (*model.Orders, error) {
	logger := klog.FromContext(ctx)

	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot start db transaction")
	}

	defer func() {
		if err != nil {
			logger.Error(err, "CreateOrder(...) transaction commit error")
			if err := tx.Rollback(); err != nil {
				logger.Error(err, "cannot rollback transaction created for CreateOrder(..)")
			}
		}
	}()

	now := time.Now()
	orderInfo := model.Orders{
		UserID:                order.UserID,
		RestaurantID:          order.RestaurantID,
		RiderID:               order.RiderID,
		TotalAmount:           order.TotalAmount,
		Status:                order.Status,
		DeliveryAddress:       order.DeliveryAddress,
		DeliveryLatitude:      order.DeliveryLatitude,
		DeliveryLongitude:     order.DeliveryLongitude,
		EstimatedDeliveryTime: order.EstimatedDeliveryTime,
		CreatedAt:             &now,
	}

	stmt := Orders.INSERT(Orders.AllColumns).MODEL(orderInfo)
	internal.LogStatement(ctx, stmt)

	result, err := stmt.ExecContext(ctx, tx)
	if err != nil {
		var mErr *mysql.MySQLError
		if errors.As(err, &mErr) {
			if mErr.Number == 1062 {
				return nil, err
			}
		}

		return nil, err
	}

	orderID, _ := result.LastInsertId()
	orderInfo.ID = orderID

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &orderInfo, nil
}

func (d *dbOperations) getMenu(ctx context.Context, restaurantId int64) ([]*model.MenuItems, error) {

	stmt := MenuItems.SELECT(MenuItems.AllColumns).
		WHERE(MenuItems.RestaurantID.EQ(Int64(restaurantId)))
	internal.LogStatement(ctx, stmt)

	var menuItems []*model.MenuItems
	if err := stmt.QueryContext(ctx, d.db, &menuItems); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	result := make([]*model.MenuItems, len(menuItems))
	copy(result, menuItems)

	return menuItems, nil
}

func (d *dbOperations) getRiderByID(ctx context.Context, riderID int64) (*model.Riders, error) {

	stmt := Riders.SELECT(Riders.AllColumns).WHERE(Riders.ID.EQ(Int64(riderID)))
	internal.LogStatement(ctx, stmt)

	var riders model.Riders
	if err := stmt.QueryContext(ctx, d.db, &riders); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		return nil, err
	}

	return &riders, nil
}

func (d *dbOperations) updateRider(ctx context.Context, riderID int64, latitude, longitude *float64) (*model.Riders, error) {
	logger := klog.FromContext(ctx)

	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot start db transaction")
	}

	defer func() {
		if err != nil {
			logger.Error(err, "updateRider(...) transaction commit error")
			if err := tx.Rollback(); err != nil {
				logger.Error(err, "cannot rollback transaction created for updateRider(..)")
			}
		}
	}()

	now := time.Now()
	riderUpdateInfo := model.Riders{
		Latitude:  latitude,
		Longitude: longitude,
		UpdatedAt: &now,
	}

	stmt := Riders.UPDATE(Riders.Latitude, Riders.Longitude, Riders.UpdatedAt).
		WHERE(Riders.ID.EQ(Int64(riderID))).
		MODEL(riderUpdateInfo)
	internal.LogStatement(ctx, stmt)

	_, err = stmt.ExecContext(ctx, d.db)
	if err != nil {
		var mErr *mysql.MySQLError
		if errors.As(err, &mErr) {
			if mErr.Number == 1062 {
				return nil, err
			}
		}

		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return d.getRiderByID(ctx, riderID)
}

func (d *dbOperations) findRestaurant(ctx context.Context, restaurant *model.Restaurants) ([]*model.Restaurants, error) {
	// Define search range (~5km)
	const latRange = 0.045
	const lngRange = 0.045

	conditions := Restaurants.IsActive.EQ(Bool(true))

	if restaurant.CuisineType != "" {
		conditions = conditions.AND(Restaurants.CuisineType.EQ(String(restaurant.CuisineType)))
	}

	if restaurant.Latitude != 0 && restaurant.Longitude != 0 {
		latMin := restaurant.Latitude - latRange
		latMax := restaurant.Latitude + latRange
		lngMin := restaurant.Longitude - lngRange
		lngMax := restaurant.Longitude + lngRange

		conditions = conditions.
			AND(Restaurants.Latitude.BETWEEN(Float(latMin), Float(latMax))).
			AND(Restaurants.Longitude.BETWEEN(Float(lngMin), Float(lngMax)))
	}

	if restaurant.PreparationTime != 0 {
		conditions = conditions.AND(Restaurants.PreparationTime.EQ(Int32(restaurant.PreparationTime)))
	}

	stmt := Restaurants.
		SELECT(Restaurants.AllColumns).
		WHERE(conditions)
	internal.LogStatement(ctx, stmt)

	var restaurants []*model.Restaurants
	if err := stmt.QueryContext(ctx, d.db, &restaurants); err != nil {
		if err.Error() == qrm.ErrNoRows.Error() {
			return nil, nil
		}
		return nil, err
	}

	return restaurants, nil
}

func (d *dbOperations) listOrders(ctx context.Context, userID, riderID, restaurantID *int64, status *biterush.Status) ([]*biterush.OrderHistory, error) {

	condition := Bool(true)

	if userID != nil {
		condition = condition.AND(Orders.UserID.EQ(Int64(*userID)))
	}
	if riderID != nil {
		condition = condition.AND(Orders.RiderID.EQ(Int64(*riderID)))
	}
	if restaurantID != nil {
		condition = condition.AND(Orders.RestaurantID.EQ(Int64(*restaurantID)))
	}
	if status != nil {
		condition = condition.AND(Orders.Status.EQ(String(string(*status))))
	}

	join := Orders.LEFT_JOIN(OrderItems, OrderItems.OrderID.EQ(Orders.ID))

	stmt := join.SELECT(Orders.AllColumns, OrderItems.AllColumns).WHERE(condition)
	internal.LogStatement(ctx, stmt)

	var rows []OrderHistoryModel
	if err := stmt.QueryContext(ctx, d.db, &rows); err != nil {
		if err.Error() == qrm.ErrNoRows.Error() {
			return nil, err
		}
		return nil, err
	}

	orderMap := make(map[int64]*biterush.OrderHistory)
	for _, r := range rows {
		if oh, exists := orderMap[r.Orders.ID]; exists {
			if r.OrderItems.ID != 0 {
				oh.OrderItems = append(oh.OrderItems, &r.OrderItems)
			}
		} else {
			oh := &biterush.OrderHistory{
				Order:      r.Orders,
				OrderItems: []*model.OrderItems{},
			}
			if r.OrderItems.ID != 0 {
				oh.OrderItems = append(oh.OrderItems, &r.OrderItems)
			}
			orderMap[r.Orders.ID] = oh
		}
	}

	result := make([]*biterush.OrderHistory, 0, len(orderMap))
	for _, oh := range orderMap {
		result = append(result, oh)
	}

	return result, nil
}

func (d *dbOperations) getOrderByID(ctx context.Context, orderID int64) (*model.Orders, error) {

	stmt := Orders.SELECT(Orders.AllColumns).WHERE(Orders.ID.EQ(Int64(orderID)))
	internal.LogStatement(ctx, stmt)

	var order model.Orders
	if err := stmt.QueryContext(ctx, d.db, &order); err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	return &order, nil
}

func (d *dbOperations) updateOrder(ctx context.Context, order *model.Orders) (*model.Orders, error) {

	now := time.Now()
	updateInfo := model.Orders{}

	if order.Status != "" {
		updateInfo.Status = order.Status
	}
	if order.RiderID != nil {
		updateInfo.RiderID = order.RiderID
	}
	updateInfo.UpdatedAt = &now

	stmt := Orders.UPDATE(Orders.Status, Orders.RiderID, Orders.UpdatedAt).
		WHERE(Orders.ID.EQ(Int64(order.ID))).
		MODEL(updateInfo)
	internal.LogStatement(ctx, stmt)

	_, err := stmt.ExecContext(ctx, d.db)
	if err != nil {
		return nil, err
	}

	return d.getOrderByID(ctx, order.ID)
}

func (d *dbOperations) getRestaurantByID(ctx context.Context, restaurantID int64) (*model.Restaurants, error) {

	stmt := Restaurants.SELECT(STAR).WHERE(Restaurants.ID.EQ(Int64(restaurantID)))
	internal.LogStatement(ctx, stmt)

	var restaurant model.Restaurants
	if err := stmt.QueryContext(ctx, d.db, &restaurant); err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, err
		}
	}

	return &restaurant, nil
}

func (d *dbOperations) findNearestAvailableRider(ctx context.Context, restaurantLng, restaurantLat float64) ([]*model.Riders, error) {

	stmt := Riders.
		SELECT(Riders.AllColumns).
		WHERE(Riders.IsAvailable.IS_TRUE()).
		ORDER_BY(
			RawFloat(
				fmt.Sprintf(
					"ST_Distance_Sphere(point(%f, %f), point(longitude, latitude))",
					restaurantLng, restaurantLat,
				),
			).ASC(),
		).
		LIMIT(5)

	var riders []*model.Riders
	if err := stmt.QueryContext(ctx, d.db, &riders); err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	if len(riders) == 0 {
		return nil, fmt.Errorf("no available riders found")
	}

	return riders, nil
}
