package todoitem

import "time"

// TodoItem model
type TodoItem struct {
	Done        bool
	Description string
	created     time.Time
	Due         time.Time
}
