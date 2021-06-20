package routers

import "github.com/gofiber/fiber/v2"

func getTodos(c *fiber.Ctx) error {
	return c.SendString("Get Todos")
}

func getTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Get Todo " + id)
}

func saveTodo(c *fiber.Ctx) error {
	return c.SendString("Save Todo")
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
