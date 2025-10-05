package restaurantlikebiz

import (
	"context"
	"crud-go/common"
	restaurantlikemodel "crud-go/module/restaurantlike/model"
)

type ListUserLikeRestaurantStore interface {
	GetUsersLikeRestaurant(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantBiz struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurantBiz(store ListUserLikeRestaurantStore) *listUserLikeRestaurantBiz {
	return &listUserLikeRestaurantBiz{store}
}

func (biz *listUserLikeRestaurantBiz) GetUsersLikeRestaurant(ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging) ([]common.SimpleUser, error) {
	data, err := biz.store.GetUsersLikeRestaurant(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}
	return data, nil
}
