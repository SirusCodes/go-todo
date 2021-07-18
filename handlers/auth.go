package handlers

import (
	"github.com/SirusCodes/go-todo/database"
	"github.com/SirusCodes/go-todo/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type AuthenticatedTokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func Register(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(Response{Status: "error", Message: "incorrect body structue"})
	}

	database.DBConn.First(&user, "username = ?", user.Username)

	if user.Username != "" {
		return c.Status(400).JSON(Response{Status: "error", Message: "username already exists"})
	}

	hash, err := hashPassword(user.Password)

	if err != nil {
		return c.Status(400).JSON(Response{Status: "error", Message: "password hashing error"})
	}

	user.Password = hash
	if err := database.DBConn.Create(&user).Error; err != nil {
		return c.Status(400).JSON(Response{Status: "error", Message: "could not create user"})
	}

	authResponse := AuthenticatedTokenResponse{}

	return c.Status(201).JSON(Response{Status: "success", Message: "successfully registered", Data: authResponse})
}

func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var loginRequest LoginRequest

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(400).JSON(Response{Status: "error", Message: "incorrect body structure"})
	}

	var user models.User
	database.DBConn.First(&user, "username = ?", loginRequest.Username)

	if user.Username == "" {
		return c.Status(403).JSON(Response{Status: "error", Message: "password or username is incorrect"})
	}

	if !verifyPassword(user.Password, loginRequest.Password) {
		return c.Status(403).JSON(Response{Status: "error", Message: "password or username is incorrect"})
	}

	authResponse := AuthenticatedTokenResponse{}

	return c.Status(200).JSON(Response{Status: "success", Message: "successfully logged in", Data: authResponse})
}
