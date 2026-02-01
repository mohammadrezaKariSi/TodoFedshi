package todo

import "time"

type ToDo struct {
	ID          uint64
	Description string
	DueDate     time.Time
	FileID      string
}
