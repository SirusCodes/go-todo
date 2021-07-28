package main

import (
	"github.com/SirusCodes/go-todo/database"
	"github.com/SirusCodes/go-todo/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	database.InitDB()

	app.Use(logger.New())
	app.Use(cors.New())

	routers.SetRouters(app)

	app.Listen(":8000")
}
