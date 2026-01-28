package main

import (
	hitrixapp "awesomeProject1/internal/infrastructure/hitrix"
	"awesomeProject1/internal/infrastructure/persistance/beeorm"
	"awesomeProject1/internal/usecase/todo"
)

func main() {
	app := hitrixapp.New()

	userRepo := beeorm.NewTodoRepository(app)
	listUsersUC := todo.NewListTodos(userRepo)

	// inject into GraphQL resolvers
	graphql.Register(app, listUsersUC)

	hitrix.Run(app)
}
