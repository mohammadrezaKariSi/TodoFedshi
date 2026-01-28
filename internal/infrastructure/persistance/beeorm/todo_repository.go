package beeorm

import (
	"awesomeProject1/internal/domain/todo"
	"context"
	"time"

	"github.com/coretrix/hitrix"
)

type ToDoEntity struct {
	ID          int64     `orm:"id"`
	Description string    `orm:"description"`
	DueDate     time.Time `orm:"due_date"`
	FileID      string    `orm:"file_id"`
}

type TodoRepository struct {
	app *hitrix.App
}

func NewTodoRepository(app *hitrix.App) *TodoRepository {
	return &TodoRepository{app: app}
}

func (r *TodoRepository) List(ctx context.Context) ([]*todo.ToDo, error) {
	orm := r.app.GetORM(ctx)

	var entities []*ToDoEntity
	if err := orm.Search(&entities); err != nil {
		return nil, err
	}

	users := make([]*todo.ToDo, 0, len(entities))
	for _, e := range entities {
		users = append(users, &todo.ToDo{
			ID:          e.ID,
			Description: e.Description,
			DueDate:     e.DueDate,
			FileID:      e.FileID,
		})
	}

	return users, nil
}
