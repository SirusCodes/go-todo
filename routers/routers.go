package routers

import (
	"github.com/SirusCodes/go-todo/handlers"
	"github.com/SirusCodes/go-todo/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetRouters(app *fiber.App) {
	app.Post("/login", handlers.Login)
	app.Post("/register", handlers.Register)
	app.Post("/refresh", handlers.GetRefreshToken)

	app.Get("/todos", middleware.HandleJWTAuth(), handlers.GetTodos)
	app.Get("/todo/:id", middleware.HandleJWTAuth(), handlers.GetTodo)
	app.Post("/todo", middleware.HandleJWTAuth(), handlers.SaveTodo)
	app.Put("/todo", middleware.HandleJWTAuth(), handlers.UpdateTodo)
	app.Delete("/todo/:id", middleware.HandleJWTAuth(), handlers.DeleteTodo)
}
