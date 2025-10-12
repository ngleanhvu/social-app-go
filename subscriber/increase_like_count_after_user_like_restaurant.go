package subscriber

import (
	"context"
	"crud-go/component/appctx"
	restaurantstorage "crud-go/module/restaurant/storage"
	"crud-go/pubsub"
	"log"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	//GetUserId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func PushNotiAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Push notification after user like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			log.Println("push notification after user like restaurant")
			return nil
		},
	}
}

func RunSomething(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Run something",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
