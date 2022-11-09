package pubsub

import (
	"fmt"
	"time"
)

type Message struct {
	id        string
	channel   string // can be ignore
	data      map[string]interface{}
	createdAt time.Time
}

func NewMessage(data map[string]interface{}) *Message {
	now := time.Now().UTC()

	return &Message{
		id:        fmt.Sprintf("%v", now.UnixNano()),
		data:      data,
		createdAt: now,
	}
}

func (evt *Message) String() string {
	return fmt.Sprintf("Message %v", evt.channel)
}
func (evt *Message) Channel() string {
	return evt.channel
}
func (evt *Message) SetChannel(channel string) {
	evt.channel = channel
}
func (evt *Message) Data() map[string]interface{} {
	return evt.data
}
