package graphql

import (
	"awesomeProject1/internal/usecase/todo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	TodoUC *todo.ListToDos
}
