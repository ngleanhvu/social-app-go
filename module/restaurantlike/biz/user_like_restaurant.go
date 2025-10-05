package restaurantlikebiz

import (
	"context"
	"crud-go/common"
	restaurantmodule "crud-go/module/restaurant/model"
	restaurantlikemodel "crud-go/module/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.RestaurantLikeCreate) error
}

type RestaurantStore interface {
	FindDataWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*restaurantmodule.Restaurant, error)
}

type userLikeRestaurantBiz struct {
	store           UserLikeRestaurantStore
	restaurantStore RestaurantStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, restaurantStore RestaurantStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, restaurantStore: restaurantStore}
}

func (biz *userLikeRestaurantBiz) UserLikeRestaurantBiz(ctx context.Context,
	data *restaurantlikemodel.RestaurantLikeCreate) error {

	restaurantData, err := biz.restaurantStore.FindDataWithCondition(ctx, map[string]interface{}{"id": data.RestaurantId})

	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodule.EntityName, err)
	}

	if restaurantData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodule.EntityName, nil)
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannonLikeRestaurant
	}

	return nil
}
