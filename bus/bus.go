package Bus

import "github.com/asaskevich/EventBus"

var bus = EventBus.New()

var Publish = bus.Publish
var Subscribe = bus.Subscribe
