package handlers

import (
	"github.com/SirusCodes/go-todo/database"
	"github.com/SirusCodes/go-todo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	database.DBConn.Find(&todos)
	return c.JSON(Response{Status: "success", Message: "Todos are listed", Data: todos})
}

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	database.DBConn.Find(&todo, "id=?", id)
	return c.JSON(Response{Status: "success", Message: "Todo found", Data: todo})
}

func SaveTodo(c *fiber.Ctx) error {

	type TodoCreateReq struct {
		Task string `validate:"required" json:"task"`
	}

	req := new(TodoCreateReq)

	if err := c.BodyParser(req); err != nil {
		return err
	}

	todo := models.Todo{
		ID:        uuid.NewString(),
		Task:      req.Task,
		Completed: false,
		CreatedBy: "USER NAME",
	}

	database.DBConn.Create(&todo)
	return c.Status(201).JSON(Response{Status: "success", Message: "Todo was created", Data: todo})
}

func UpdateTodo(c *fiber.Ctx) error {

	type TodoUpdateReq struct {
		Task      string `json:"task"`
		ID        string `json:"id"`
		Completed bool   `json:"completed"`
	}

	req := new(TodoUpdateReq)

	if err := c.BodyParser(req); err != nil {
		return err
	}

	todo := models.Todo{Completed: req.Completed, Task: req.Task, ID: req.ID}
	database.DBConn.Model(&todo).Update("completed", "task")

	return c.JSON(Response{Status: "success", Message: "Todo was updated", Data: todo})
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DBConn.Delete(&models.Todo{}, "id=?", id)
	return c.Status(204).JSON(Response{Status: "success", Message: "Todo was deleted"})
}
