package beeorm

import (
	"awesomeProject1/internal/domain/todo"
	"context"
	"time"
	//"github.com/latolukasz/beeorm"
)

type ToDoEntity struct {
	ID          int64     `orm:"id"`
	Description string    `orm:"description"`
	DueDate     time.Time `orm:"due_date"`
	FileID      string    `orm:"file_id"`
}

func (*ToDoEntity) TableName() string {
	return "todos"
}

//
//func Init(registry *beeorm.Registry) {
//	registry.RegisterEntity(&ToDoEntity{})
//}

type TodoRepository struct{}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) List(ctx context.Context) ([]*todo.ToDo, error) {
	//o := orm.MustFromContext(ctx)
	//
	//var entities []*ToDoEntity
	//if err := o.Find(&entities); err != nil {
	//	return nil, err
	//}
	//
	//todos := make([]*todo.ToDo, 0, len(entities))
	//for _, e := range entities {
	//	todos = append(todos, &todo.ToDo{
	//		ID:          e.ID,
	//		Description: e.Description,
	//		DueDate:     e.DueDate,
	//		FileID:      e.FileID,
	//	})
	//}
	//
	//return todos, nil

	return nil, nil
}
