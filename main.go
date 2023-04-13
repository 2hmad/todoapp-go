package main

import (
	"context"
	"github.com/gopulse/pulse"
	"github.com/todoapp-pulse/migrations"
	"github.com/todoapp-pulse/todoRoutes"
)

func main() {
	app := pulse.New()
	router := pulse.NewRouter()

	app.Router = router

	router.Get("/", func(c *pulse.Context) error {
		c.String("Hello")
		return nil
	})

	// Todo routes
	todoRoutes.TodoRoute(router)

	router.Use("GET", pulse.CORSMiddleware())

	// Run migrations
	migrations.M001CreateTodosTable(context.Background())

	app.Run(":8085")
}
