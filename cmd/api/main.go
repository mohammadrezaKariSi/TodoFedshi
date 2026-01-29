package main

import (
	"awesomeProject1/internal/delivery/graphql"
	hitrixapp "awesomeProject1/internal/infrastructure/hitrix"
	beeinfra "awesomeProject1/internal/infrastructure/persistence/beeorm"
	"awesomeProject1/internal/usecase/todo"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/latolukasz/beeorm"
)

func main() {
	app, cleanup := hitrixapp.New()
	defer cleanup()

	app.RegisterService(&beeorm.MySQLPool{
		DSN:          "todo_user:todo_pass@tcp(localhost:3306)/tododb?charset=utf8mb4&parseTime=True&loc=UTC",
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		MaxLifetime:  time.Hour,
		Alias:        "default",
	})

	todoRepo := beeinfra.NewTodoRepository()
	todoUC := todo.NewListTodos(todoRepo)

	resolver := &graphql.Resolver{
		TodoUC: todoUC,
	}

	executableSchema := graphql.NewExecutableSchema(
		graphql.Config{Resolvers: resolver},
	)

	gqlHandler := handler.NewDefaultServer(executableSchema)

	app.RunServer(
		8080,
		func(ginEngine *gin.Engine) {
			ginEngine.POST("/graphql", func(c *gin.Context) {
				gqlHandler.ServeHTTP(c.Writer, c.Request)
			})

			ginEngine.GET("/playql", func(c *gin.Context) {
				playground.Handler("GraphQL", "/graphql").ServeHTTP(c.Writer, c.Request)
			})

			ginEngine.GET("/health", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
				})
			})

		},
	)
}
