package todo

import "context"

type Repository interface {
	List(ctx context.Context) ([]*ToDo, error)
}
