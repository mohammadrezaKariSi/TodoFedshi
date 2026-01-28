package todo

import (
	domain "awesomeProject1/internal/domain/todo"
	"context"
)

type ListToDos struct {
	repo domain.Repository
}

func NewListTodos(repo domain.Repository) *ListToDos {
	return &ListToDos{repo: repo}
}

func (uc *ListToDos) Execute(ctx context.Context) ([]*domain.ToDo, error) {
	return uc.repo.List(ctx)
}
