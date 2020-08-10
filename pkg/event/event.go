//go:generate stringer -type=Event

package event

// Event is operation event
type Event int

const (
	// EventCreate is create entity
	EventCreate Event = iota + 1
	// EventUpdate is update entity
	EventUpdate
	// EventDelete is delete entity
	EventDelete
)
