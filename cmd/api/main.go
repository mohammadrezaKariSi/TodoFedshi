package main

import (
	"awesomeProject1/internal/delivery/graphql"
	hitrixapp "awesomeProject1/internal/infrastructure/hitrix"
	beeinfra "awesomeProject1/internal/infrastructure/persistence/beeorm"
	"awesomeProject1/internal/scripts"
	"awesomeProject1/internal/usecase/todo"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/coretrix/hitrix"
	"github.com/coretrix/hitrix/service"
	"github.com/coretrix/hitrix/service/registry"
	"github.com/gin-gonic/gin"
)

func main() {
	app, cleanup := hitrixapp.New()
	defer cleanup()

	go app.RunBackgroundProcess(func(b *hitrix.BackgroundProcessor) {
		b.RunScript(&scripts.TodoScript{})
	})

	registry.ServiceProviderConfigDirectory("./config")

	service.DI().Config()

	var ormEngine = service.DI().OrmEngine()

	alters := ormEngine.GetAlters()

	for _, alter := range alters {
		fmt.Printf("🛠 Applying change: %s\n", alter.SQL)
		alter.Exec()
	}

	todoRepo := beeinfra.NewTodoRepository(ormEngine)
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
