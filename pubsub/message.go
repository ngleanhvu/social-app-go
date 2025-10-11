package pubsub

import (
	"fmt"
	"time"
)

type Message struct {
	id      string
	channel Topic // can be ignore
	data    interface{}
	time    time.Time
}

func NewMessage(data interface{}) *Message {
	return &Message{
		id:   fmt.Sprintf("%d\n", time.Now().Unix()),
		data: data,
		time: time.Now(),
	}
}

func (evt *Message) String() string {
	return fmt.Sprintf("Message %s", evt.channel)
}

func (evt *Message) Channel() Topic {
	return evt.channel
}

func (evt *Message) Data() interface{} {
	return evt.data
}

func (evt *Message) SetChannel(channel Topic) {
	evt.channel = channel
}
