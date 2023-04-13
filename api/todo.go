package api

import (
	"github.com/gopulse/pulse"
	"github.com/todoapp-pulse/controllers"
)

func TodoRoute(router *pulse.Router) {
	api := &pulse.Group{
		Prefix: "/api",
		Router: router,
	}

	v1 := api.Group("/v1")

	v1.GET("/todos", func(c *pulse.Context) error {
		_, err := controllers.GetTodos(c)
		if err != nil {
			return err
		}
		return nil
	})

	v1.POST("/todos", func(c *pulse.Context) error {
		_, err := controllers.PostTodos(c)
		c.SetRequestHeader("Access-Control-Allow-Origin", "*")
		c.SetRequestHeader("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.SetRequestHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if err != nil {
			return err
		}
		return nil
	})

	v1.PUT("/todos/:id", func(c *pulse.Context) error {
		_, err := controllers.PutTodos(c)
		if err != nil {
			return err
		}
		return nil
	})

}
