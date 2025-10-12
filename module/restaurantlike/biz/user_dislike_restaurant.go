package restaurantlikebiz

import (
	"context"
	"crud-go/common"
	restaurantmodule "crud-go/module/restaurant/model"
	restaurantlikemodel "crud-go/module/restaurantlike/model"
	"crud-go/pubsub"
	"log"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, data *restaurantlikemodel.RestaurantLikeUpdate) error
}

//type UserDislikeRestaurantDecreaseStore interface {
//	DecreaseLikeCount(ctx context.Context, id int) error
//}

type userDislikeRestaurantBiz struct {
	store           UserDislikeRestaurantStore
	restaurantStore RestaurantStore
	//decreaseLikeCount UserDislikeRestaurantDecreaseStore
	pb pubsub.PubSub
}

func NewDislikeRestaurantBiz(store UserDislikeRestaurantStore,
	restaurantStore RestaurantStore,
	pb pubsub.PubSub) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store,
		restaurantStore: restaurantStore,
		//decreaseLikeCount: decreaseLikeCount,
		pb: pb,
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
	log.Println("data", data)
	if err := biz.pb.Publish(ctx, common.TopicUserDislikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println("Error in dislike", err)
	}

	// Side effect
	//j := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.decreaseLikeCount.DecreaseLikeCount(ctx, data.RestaurantId)
	//})
	//
	//if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	//	log.Println(err)
	//}

	return nil
}
