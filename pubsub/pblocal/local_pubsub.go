package pblocal

import (
	"context"
	"crud-go/common"
	"crud-go/pubsub"
	"log"
	"sync"
)

type localPubSub struct {
	messageQueue chan *pubsub.Message
	mapChanel    map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewLocalPubSub() *localPubSub {
	pb := &localPubSub{
		messageQueue: make(chan *pubsub.Message, 10000),
		mapChanel:    make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       &sync.RWMutex{},
	}
	pb.run()
	return pb
}

func (pb *localPubSub) Publish(ctx context.Context,
	topic pubsub.Topic,
	data *pubsub.Message) error {
	data.SetChannel(topic)

	go func() {
		defer common.AppRecover()
		pb.messageQueue <- data
		log.Println("New event published", data.String(), " with data", data.Data())
	}()
	return nil
}

func (pb *localPubSub) Subscribe(ctx context.Context,
	topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)
	pb.locker.Lock()

	if val, ok := pb.mapChanel[topic]; ok {
		val = append(pb.mapChanel[topic], c)
		pb.mapChanel[topic] = val
	} else {
		pb.mapChanel[topic] = []chan *pubsub.Message{c}
	}
	pb.locker.Unlock()

	return c, func() {
		log.Println("Close")

		if chans, ok := pb.mapChanel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					chans = append(chans[:i], chans[i+1:]...)
					pb.locker.Lock()
					pb.mapChanel[topic] = chans
					pb.locker.Unlock()
					break
				}
			}
		}
	}
}

func (pb *localPubSub) run() {
	log.Println("Running pubsub")
	go func() {
		defer common.AppRecover()
		for {
			message := <-pb.messageQueue
			log.Print("Message dequeue", message)
			if subs, ok := pb.mapChanel[message.Channel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						defer common.AppRecover()
						c <- message
					}(subs[i])
				}
			}
		}
	}()
}
