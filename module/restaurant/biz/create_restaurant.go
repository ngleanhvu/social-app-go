package restaurantbiz

import (
	"context"
	"crud-go/common"
	"crud-go/module/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(context context.Context, data *restaurantmodule.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodule.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodule.EntityName, err)
	}
	return nil
}
