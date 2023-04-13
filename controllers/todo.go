package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gopulse/pulse"
	"github.com/todoapp-pulse/config"
	"github.com/todoapp-pulse/models"
)

func GetTodos(c *pulse.Context) ([]byte, error) {
	db, err := config.Connect()

	var todos []models.Todo

	err = db.Select(&todos, "SELECT * FROM todos")
	if err != nil {
		return nil, err
	}

	return c.JSON(200, todos)
}

func PostTodos(c *pulse.Context) ([]byte, error) {
	db, err := config.Connect()

	var todo models.Todo

	err = c.BodyParser(&todo)
	if err != nil {
		return nil, err
	}

	_, err = db.NamedExec("INSERT INTO todos (title, done, created_at, updated_at) VALUES (:title, :done, now(), now())", todo)
	if err != nil {
		return nil, err
	}

	return c.JSON(200, todo)
}

func PutTodos(c *pulse.Context) ([]byte, error) {
	db, err := config.Connect()

	var todo models.Todo

	err = c.BodyParser(&todo)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("UPDATE todos SET title = ?, done = ?, updated_at = now() WHERE id = ?", todo.Title, todo.Done, c.Param("id"))
	if err != nil {
		return nil, err
	}

	return c.JSON(200, todo)
}
