package main

import (
	"context"
	"crud-go/pubsub"
	"crud-go/pubsub/pblocal"
	"log"
	"time"
)

func main() {
	localPubSub := pblocal.NewLocalPubSub()

	var topic pubsub.Topic = "OrderCreated"

	sub1, _ := localPubSub.Subscribe(context.Background(), topic)
	sub2, _ := localPubSub.Subscribe(context.Background(), topic)

	localPubSub.Publish(context.Background(), topic, pubsub.NewMessage(1))
	localPubSub.Publish(context.Background(), topic, pubsub.NewMessage(2))

	go func() {
		for {
			log.Println("Sub 1", (<-sub1).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	go func() {
		for {
			log.Println("Sub 2", (<-sub2).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	time.Sleep(time.Second * 3)

}
