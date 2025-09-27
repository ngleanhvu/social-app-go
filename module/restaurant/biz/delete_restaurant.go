package restaurantbiz

import (
	"context"
	"crud-go/module/restaurant/model"
	"errors"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodule.Restaurant, error)
	Delete(context context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (b *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {

	oldData, err := b.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("restaurant has been deleted")
	}

	if err := b.store.Delete(context, id); err != nil {
		return err
	}

	return nil
}
