package beeorm

import (
	"awesomeProject1/internal/domain/todo"
	"context"
	"time"

	"github.com/latolukasz/beeorm"
)

const (
	RedisPool   = "redis_pool"
	ToDoChannel = "todo_channel"
	ToDoGroup   = "todo_channel_group"
)

type ToDoEntity struct {
	beeorm.ORM  `orm:"table=todos;dirty=de"`
	ID          uint64    `orm:"id"`
	Description string    `orm:"description"`
	DueDate     time.Time `orm:"due_date"`
	FileID      string    `orm:"file_id"`
}

func Init(registry *beeorm.Registry) {
	registry.RegisterEntity(
		&ToDoEntity{},
	)

	registry.RegisterRedisStream(ToDoChannel, RedisPool, []string{ToDoGroup})
}

type TodoRepository struct {
	engine *beeorm.Engine
}

func NewTodoRepository(engine *beeorm.Engine) *TodoRepository {
	return &TodoRepository{engine: engine}
}

func (r *TodoRepository) List(ctx context.Context) ([]*todo.ToDo, error) {
	var entities []*ToDoEntity

	where := beeorm.NewWhere("1")

	r.engine.Search(where, nil, &entities)

	todos := make([]*todo.ToDo, 0, len(entities))
	for _, e := range entities {
		todos = append(todos, &todo.ToDo{
			ID:          e.ID,
			Description: e.Description,
			DueDate:     e.DueDate,
			FileID:      e.FileID,
		})
	}

	return todos, nil
}

func (r *TodoRepository) Create(ctx context.Context, inp *todo.ToDo) (*todo.ToDo, error) {
	entity := &ToDoEntity{
		Description: inp.Description,
		FileID:      inp.FileID,
		DueDate:     inp.DueDate,
	}

	r.engine.Flush(entity)

	return &todo.ToDo{
		ID:          entity.ID,
		Description: entity.Description,
		DueDate:     entity.DueDate,
		FileID:      entity.FileID,
	}, nil
}
