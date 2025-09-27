package restaurantbiz

import (
	"context"
	"crud-go/common"
	restaurantmodule "crud-go/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(context context.Context,
		filter *restaurantmodule.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodule.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurantBiz(ctx context.Context,
	filter *restaurantmodule.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodule.Restaurant, error) {

	result, err := biz.store.ListDataWithCondition(ctx, filter, paging, moreKeys...)

	if err != nil {
		return nil, err
	}

	return result, nil

}
