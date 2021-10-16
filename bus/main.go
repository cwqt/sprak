package Bus

import (
	eb "github.com/asaskevich/EventBus"
)

var bus = eb.New()

type Event struct {
	Topic string
	Data  interface{}
}

func Publish(topic string, args ...interface{}) {
	bus.Publish(topic, Event{Topic: topic, Data: args})
}

var Subscribe = bus.Subscribe
