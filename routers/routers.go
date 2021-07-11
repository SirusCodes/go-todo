package routers

import (
	"github.com/SirusCodes/go-todo/database"
	"github.com/SirusCodes/go-todo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	database.DBConn.Find(&todos)
	return c.JSON(todos)
}

func getTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Get Todo " + id)
}

func saveTodo(c *fiber.Ctx) error {
	todo := models.Todo{
		ID:        uuid.UUID{},
		Task:      "Task",
		Completed: false,
		CreatedBy: "Darshan",
	}
	database.DBConn.Create(&todo)
	return c.JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("UpdateTodo " + id)
}

func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("DeleteTodo " + id)
}

func SetRouters(app *fiber.App) {
	app.Get("/todos", getTodos)
	app.Get("/todo/:id", getTodo)
	app.Post("/todo", saveTodo)
	app.Put("/todo/:id", updateTodo)
	app.Delete("/todo/:id", deleteTodo)
}
