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

func (uc *ListToDos) CreateTodo(ctx context.Context, inp *domain.ToDo) (*domain.ToDo, error) {
	todo, err := uc.repo.Create(ctx, inp)
	return todo, err
}

func (uc *ListToDos) Execute(ctx context.Context) ([]*domain.ToDo, error) {
	return uc.repo.List(ctx)
}
