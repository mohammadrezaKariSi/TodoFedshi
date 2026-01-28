package todo

import "time"

type ToDo struct {
	ID          int64
	Description string
	DueDate     time.Time
	FileID      string
}
