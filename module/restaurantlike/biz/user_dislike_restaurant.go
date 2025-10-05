package restaurantlikebiz

import (
	"context"
	"crud-go/common"
	restaurantmodule "crud-go/module/restaurant/model"
	restaurantlikemodel "crud-go/module/restaurantlike/model"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, data *restaurantlikemodel.RestaurantLikeUpdate) error
}

type userDislikeRestaurantBiz struct {
	store           UserDislikeRestaurantStore
	restaurantStore RestaurantStore
}

func NewDislikeRestaurantBiz(store UserDislikeRestaurantStore,
	restaurantStore RestaurantStore) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store,
		restaurantStore: restaurantStore,
	}
}

func (biz *userDislikeRestaurantBiz) UserDislikeRestaurantBiz(ctx context.Context,
	data *restaurantlikemodel.RestaurantLikeUpdate) error {

	restaurantData, err := biz.restaurantStore.FindDataWithCondition(ctx,
		map[string]interface{}{"id": data.RestaurantId})

	if restaurantData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodule.EntityName, err)
	}

	if err := biz.store.Delete(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannonDislikeRestaurant
	}

	return nil
}
