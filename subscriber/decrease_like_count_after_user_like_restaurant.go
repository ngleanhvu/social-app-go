package subscriber

import (
	"context"
	"crud-go/component/appctx"
	restaurantstorage "crud-go/module/restaurant/storage"
	"crud-go/pubsub"
)

//func DecreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext,
//	ctx context.Context) {
//	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserDislikeRestaurant)
//	store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
//	go func() {
//		defer common.AppRecover()
//		for {
//			msg := <-c
//			likeData := msg.Data().(HasRestaurantId)
//			_ = store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
//		}
//	}()
//}

func DecreaseLikeCountAfterUserLikeRestaurant(aptCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after use dislike restaurant!",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSqlStore(aptCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
