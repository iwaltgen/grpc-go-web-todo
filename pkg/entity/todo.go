package entity

import "time"

// Todo entity
type Todo struct {
	ID          string
	Description string
	Completed   bool
	ModifiedAt  time.Time
	CreatedAt   time.Time
}
