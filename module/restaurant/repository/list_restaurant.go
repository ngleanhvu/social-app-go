package restaurantrepository

import (
	"context"
	"crud-go/common"
	restaurantmodule "crud-go/module/restaurant/model"
	"log"
)

type ListRestaurantStore interface {
	ListDataWithCondition(ctx context.Context,
		filter *restaurantmodule.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodule.Restaurant, error)
}

type LikeRestaurantStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantRepo struct {
	store     ListRestaurantStore
	likeStore LikeRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore, likeStore LikeRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{store, likeStore}
}

func (repo *listRestaurantRepo) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodule.Filter,
	paging *common.Paging) ([]restaurantmodule.Restaurant, error) {
	result, err := repo.store.ListDataWithCondition(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodule.EntityName, err)
	}

	ids := make([]int, len(result))

	for k, v := range result {
		ids[k] = v.Id
	}

	likeMap, err := repo.likeStore.GetRestaurantLikes(ctx, ids)

	if err != nil {
		log.Println(err)
		return result, nil
	}

	for k, v := range result {
		result[k].LikeCount = likeMap[v.Id]
	}

	return result, nil
}
