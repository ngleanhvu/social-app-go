package restaurantbiz

import (
	"context"
	"crud-go/common"
	"crud-go/module/restaurant/model"
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
		return common.ErrEntityNotFound(restaurantmodule.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodule.EntityName, nil)
	}

	if err := b.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodule.EntityName, err)
	}

	return nil
}
