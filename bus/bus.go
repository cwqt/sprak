package Bus

import (
	hub "github.com/simonfxr/pubsub"
)

var bus = hub.NewBus()

type Event struct {
	Topic string
	Data  interface{}
}

func Publish(topic string, data interface{}) {
	bus.Publish(topic, Event{Topic: topic, Data: data})
}

func Subscribe(topic string, cb func(event Event)) func() {
	sub := bus.SubscribeAsync(topic, cb)

	return func() {
		// provide utility for un-subscribing
		bus.Unsubscribe(sub)
	}
}
