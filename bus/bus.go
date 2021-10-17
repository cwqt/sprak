package Bus

import (
	tea "github.com/charmbracelet/bubbletea"
	hub "github.com/simonfxr/pubsub"
)

var bus = hub.NewBus()
var program *tea.Program

type Event struct {
	Topic string
	Data  interface{}
}

func AttachToProgram(p *tea.Program) {
	program = p
}

func Publish(topic string, data interface{}) {
	// Not entirely sure why this needs to be in a goroutine
	// but otherwise the entire program gets locked up
	go func() {
		event := Event{Topic: topic, Data: data}
		bus.Publish(topic, event)

		// fmt.Printf("%+v\n", program)
		if program != nil {
			program.Send(event)
		}
	}()
}

func Subscribe(topic string, cb func(event Event)) func() {
	sub := bus.SubscribeAsync(topic, cb)

	return func() {
		// provide utility for un-subscribing
		bus.Unsubscribe(sub)
	}
}
