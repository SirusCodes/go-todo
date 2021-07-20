package handlers

import (
	"github.com/SirusCodes/go-todo/database"
	"github.com/SirusCodes/go-todo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func getUserFromToken(c *fiber.Ctx) string {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	return claims["username"].(string)
}

func GetTodos(c *fiber.Ctx) error {
	uid := getUserFromToken(c)
	var todos []models.Todo
	database.DBConn.Find(&todos, "created_by=?", uid)
	return c.JSON(Response{Status: "success", Message: "Todos are listed", Data: todos})
}

func GetTodo(c *fiber.Ctx) error {
	uid := getUserFromToken(c)
	id := c.Params("id")
	var todo models.Todo
	database.DBConn.Find(&todo, "id=? AND created_by=?", id, uid)
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

	uid := getUserFromToken(c)

	todo := models.Todo{
		ID:        uuid.NewString(),
		Task:      req.Task,
		Completed: false,
		CreatedBy: uid,
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

	uid := getUserFromToken(c)

	todo := models.Todo{Completed: req.Completed, Task: req.Task, ID: req.ID, CreatedBy: uid}
	database.DBConn.Model(&todo).Update("completed", "task")

	return c.JSON(Response{Status: "success", Message: "Todo was updated", Data: todo})
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	uid := getUserFromToken(c)
	database.DBConn.Delete(&models.Todo{}, "id=? AND created_by=?", id, uid)
	return c.Status(204).JSON(Response{Status: "success", Message: "Todo was deleted"})
}
