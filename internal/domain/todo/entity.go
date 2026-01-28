package todo

import "time"

type ToDo struct {
	ID          int64     `orm:"id"`
	Description string    `orm:"description"`
	DueDate     time.Time `orm:"due_date"`
	FileID      string    `orm:"file_id"`
}
