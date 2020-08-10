package event

import (
	"sync"

	"go.uber.org/atomic"
)

type (
	// Option is subscribe option
	Option func(event Event) bool
	// Subscription is subscribe handler
	Subscription func(event Event, value interface{})
)

// Emitter event emitter
type Emitter struct {
	generator     atomic.Uint32
	subscriptions sync.Map
}

// Subscribe is subscribe event
func (e *Emitter) Subscribe(subscription Subscription, option Option) func() {
	id := e.generator.Inc()
	e.subscriptions.Store(id, e.wrap(subscription, option))
	return func() {
		e.subscriptions.Delete(id)
	}
}

// Publish 이벤트 발행
func (e *Emitter) Publish(event Event, value interface{}) {
	e.subscriptions.Range(func(_, fn interface{}) bool {
		go fn.(Subscription)(event, value)
		return true
	})
}

func (e *Emitter) wrap(subscription Subscription, option Option) Subscription {
	return func(event Event, value interface{}) {
		if option(event) {
			subscription(event, value)
		}
	}
}

// WithEvent is subscribe specific event
func WithEvent(evt Event) Option {
	return func(event Event) bool {
		return event == evt
	}
}
