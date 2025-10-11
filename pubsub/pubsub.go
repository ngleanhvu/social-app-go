package pubsub

import "context"

type Topic string

type PubSub interface {
	Publish(ctx context.Context, chanel Topic, data *Message) error
	Subscribe(ctx context.Context, chanel Topic) (ch <-chan *Message, close func()) // close để unsubscribe
}
