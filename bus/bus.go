package Bus

import eb "github.com/asaskevich/EventBus"

var bus = eb.New()

type Event struct {
	Topic string
	Data  interface{}
}

func Publish(topic string, data interface{}) {
	bus.Publish(topic, Event{Topic: topic, Data: data})
}

func Subscribe(topic string, fn interface{}) {
	bus.Subscribe(topic, fn)
}
