package routers

import (
	"github.com/SirusCodes/go-todo/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetRouters(app *fiber.App) {
	app.Get("/todos", handlers.GetTodos)
	app.Get("/todo/:id", handlers.GetTodo)
	app.Post("/todo", handlers.SaveTodo)
	app.Put("/todo", handlers.UpdateTodo)
	app.Delete("/todo/:id", handlers.DeleteTodo)
}
