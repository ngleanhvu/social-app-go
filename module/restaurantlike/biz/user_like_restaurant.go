package restaurantlikebiz

import (
	"context"
	"crud-go/common"
	restaurantmodule "crud-go/module/restaurant/model"
	restaurantlikemodel "crud-go/module/restaurantlike/model"
	"crud-go/pubsub"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.RestaurantLikeCreate) error
}

//type IncreaseLikeCountStore interface {
//	IncreaseLikeCount(ctx context.Context, id int) error
//}

type RestaurantStore interface {
	FindDataWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*restaurantmodule.Restaurant, error)
}

type userLikeRestaurantBiz struct {
	store           UserLikeRestaurantStore
	restaurantStore RestaurantStore
	//increaseLikeCountStore IncreaseLikeCountStore
	pb pubsub.PubSub
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore,
	restaurantStore RestaurantStore,
	//increaseLikeCountStore IncreaseLikeCountStore,
	pb pubsub.PubSub) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store,
		restaurantStore: restaurantStore,
		//increaseLikeCountStore: increaseLikeCountStore,
		pb: pb,
	}
}

func (biz *userLikeRestaurantBiz) UserLikeRestaurantBiz(ctx context.Context,
	data *restaurantlikemodel.RestaurantLikeCreate) error {

	restaurantData, _ := biz.restaurantStore.FindDataWithCondition(ctx, map[string]interface{}{"id": data.RestaurantId})

	if restaurantData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodule.EntityName, nil)
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannonLikeRestaurant
	}

	biz.pb.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))

	//j := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.increaseLikeCountStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//})
	//
	//if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	//	log.Println(err)
	//}

	return nil
}
