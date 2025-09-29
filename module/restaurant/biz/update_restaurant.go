package restaurantbiz

import (
	"context"
	"crud-go/common"
	"crud-go/module/restaurant/model"
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
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodule.EntityName, err)
	}
	return nil
}
