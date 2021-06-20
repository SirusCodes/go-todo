package main

import (
	"github.com/SirusCodes/go-todo/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routers.SetRouters(app)

	app.Listen(":8000")
}
