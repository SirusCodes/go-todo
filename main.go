package main

import (
	"github.com/SirusCodes/go-todo/database"
	"github.com/SirusCodes/go-todo/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.InitDB()

	routers.SetRouters(app)

	app.Listen(":8000")
}
