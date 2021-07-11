package routers

import (
	"github.com/SirusCodes/go-todo/database"
	"github.com/SirusCodes/go-todo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/go-playground/validator.v9"
)

type TodoCreateReq struct {
	Task      string `validate:"required"`
	CreatedBy string `validate:"required"`
}

type TodoUpdateReq struct {
	Task      string
	ID        string
	Completed bool
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func validateTodoCreatedReq(todoReq TodoCreateReq) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(todoReq)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func getTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	database.DBConn.Find(&todos)
	return c.JSON(todos)
}

func getTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	database.DBConn.Find(&todo, "id=?", id)
	return c.JSON(todo)
}

func saveTodo(c *fiber.Ctx) error {
	req := new(TodoCreateReq)

	if err := c.BodyParser(req); err != nil {
		return err
	}

	if err := validateTodoCreatedReq(*req); err != nil {
		return c.JSON(err)
	}

	todo := models.Todo{
		ID:        uuid.NewString(),
		Task:      req.Task,
		Completed: false,
		CreatedBy: req.CreatedBy,
	}

	database.DBConn.Create(&todo)
	return c.Status(201).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	req := new(TodoUpdateReq)

	if err := c.BodyParser(req); err != nil {
		return err
	}

	database.DBConn.Model(&models.Todo{}).Where("id=?", req.ID).Updates(models.Todo{Completed: req.Completed, Task: req.Task})

	return c.SendString("Todo was updated")
}

func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DBConn.Delete(&models.Todo{}, "id=?", id)
	return c.Status(202).SendString("Todo was deleted")
}

func SetRouters(app *fiber.App) {
	app.Get("/todos", getTodos)
	app.Get("/todo/:id", getTodo)
	app.Post("/todo", saveTodo)
	app.Put("/todo", updateTodo)
	app.Delete("/todo/:id", deleteTodo)
}
