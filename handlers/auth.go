package handlers

import (
	"errors"
	"time"

	"github.com/SirusCodes/go-todo/config"
	"github.com/SirusCodes/go-todo/database"
	"github.com/SirusCodes/go-todo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticatedTokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Username     string `json:"username"`
}

type JWTPayload struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Expires  int64  `json:"exp"`
}

func generateJWTToken(username string, sign []byte, days time.Duration) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTPayload{
		Username: username,
		Expires:  time.Now().Add(time.Hour * 24 * days).Unix(),
	})
	t, err := token.SignedString(sign)

	if err != nil {
		return "", err
	}
	return t, nil
}

func getTokenResponse(username string) (AuthenticatedTokenResponse, error) {
	token, err := generateJWTToken(username, []byte(config.Config("TOKEN")), 1)
	if err != nil {
		return AuthenticatedTokenResponse{}, err
	}

	refreshToken, err := generateJWTToken(username, []byte(config.Config("REFRESH_TOKEN")), 5)
	if err != nil {
		return AuthenticatedTokenResponse{}, err
	}

	return AuthenticatedTokenResponse{Token: token, RefreshToken: refreshToken, Username: username}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func verifyToken(token string, sign string) (*JWTPayload, error) {
	tkn, err := jwt.ParseWithClaims(token, &JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config(sign)), nil
	})

	if err != nil {
		return nil, err
	}

	payload, ok := tkn.Claims.(*JWTPayload)
	if !ok {
		return nil, errors.New("cannot parse claims")
	}

	return payload, nil
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

	authResponse, err := getTokenResponse(user.Username)
	if err != nil {
		return c.Status(400).JSON(Response{Status: "error", Message: "could not generate token"})
	}

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

	if user.Username == "" || !verifyPassword(user.Password, loginRequest.Password) {
		return c.Status(403).JSON(Response{Status: "error", Message: "password or username is incorrect"})
	}

	authResponse, err := getTokenResponse(user.Username)
	if err != nil {
		return c.Status(400).JSON(Response{Status: "error", Message: "could not generate token"})
	}

	return c.Status(200).JSON(Response{Status: "success", Message: "successfully logged in", Data: authResponse})
}

func GetRefreshToken(c *fiber.Ctx) error {
	type RefreshTokenRequest struct {
		Username     string `json:"username"`
		RefreshToken string `json:"refresh_token"`
	}
	var refreshTokenRequest RefreshTokenRequest

	if err := c.BodyParser(&refreshTokenRequest); err != nil {
		return c.Status(400).JSON(Response{Status: "error", Message: "incorrect body structure"})
	}

	token, err := verifyToken(refreshTokenRequest.RefreshToken, "REFRESH_TOKEN")
	if err != nil || token.Username != refreshTokenRequest.Username {
		return c.Status(400).JSON(Response{Status: "error", Message: "invalid refresh token"})
	}

	newToken, err := generateJWTToken(refreshTokenRequest.Username, []byte(config.Config("TOKEN")), 1)

	if err != nil {
		return c.Status(400).JSON(Response{Status: "error", Message: "could not generate token"})
	}

	return c.Status(201).JSON(Response{Status: "success", Message: "successfully refreshed token", Data: fiber.Map{
		"token": newToken,
	}})
}
