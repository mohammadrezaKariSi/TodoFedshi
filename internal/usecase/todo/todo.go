package todo

import (
	domain "awesomeProject1/internal/domain/todo"
	"context"
)

type ListUsers struct {
	repo domain.Repository
}

func NewListUsers(repo domain.Repository) *ListUsers {
	return &ListUsers{repo: repo}
}

func (uc *ListUsers) Execute(ctx context.Context) ([]*domain.ToDo, error) {
	return uc.repo.List(ctx)
}
