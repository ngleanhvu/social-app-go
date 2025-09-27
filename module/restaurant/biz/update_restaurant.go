package restaurantbiz

import (
	"context"
	"crud-go/module/restaurant/model"
	"errors"
)

// Declare a interface to interact with storage layer because if not use interface
// you must declare implementation of storage layer in biz layer
// => You will not obtain loose coupling and clean architectur

type UpdateRestaurantStore interface {
	Update(context context.Context, id int, data *restaurantmodule.RestaurantUpdate) error
}
type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, data *restaurantmodule.RestaurantUpdate, id int) error {
	if data.Name != nil && *data.Name == "" {
		return errors.New("restaurant name is required")
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return err
	}
	return nil
}
