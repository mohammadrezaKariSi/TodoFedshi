package todo

import "context"

type Repository interface {
	List(ctx context.Context) ([]*ToDo, error)
	Create(ctx context.Context, do *ToDo) (*ToDo, error)
}
