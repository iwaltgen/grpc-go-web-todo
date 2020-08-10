package message

import (
	"github.com/iwaltgen/grpc-go-web-todo/pkg/event"

	todov1 "github.com/iwaltgen/grpc-go-web-todo/api/todo/v1"
)

// EventFromProto convert todov1.Event to event.Event
func EventFromProto(value todov1.Event) event.Event {
	return event.Event(value)
}

// EventProto convert event.Event to todov1.Event
func EventProto(value event.Event) todov1.Event {
	return todov1.Event(value)
}
