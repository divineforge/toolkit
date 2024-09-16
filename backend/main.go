package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/swaggo/swag/cmd/swag"
)

import "github.com/arsmn/fiber-swagger/swagger"

// @title Fiber Swagger Example API
// @version 1.0
// @description This is a sample API using Fiber and Swagger.
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Swagger setup
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// API endpoints
	app.Get("/users", getUsers)
	app.Get("/users/:id", getUser)
	app.Post("/users", createUser)
	app.Put("/users/:id", update)
	app.Delete("/users/:id", deleteUser)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// @Summary Get all users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(c *fiber.Ctx) error {
	users := []User{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Doe"},
	}
	return c.JSON(users)
}

// @Summary Get user by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
func getUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("Invalid user ID")
	}
	user := User{ID: id, Name: "John Doe"}
	return c.JSON(user)
}

// @Summary Create new user
// @Produce json
// @Param user body User true "User data"
// @Success 201 {object} User
// @Router /users [post]
func createUser(c *fiber.Ctx) error {
	var user User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).SendString("Invalid user data")
	}
	return c.JSON(user)
}

// @Summary Update existing user
// @Produce json
// @Param id path int true "User ID"
// @Param user body User true "User data"
// @Success 200 {object} User
// @Router /users/{id} [put]
func update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("Invalid user ID")
	}
	var user User
	err = c.BodyParser(&user)
	if err != nil {
		return c.Status(400).SendString("Invalid user data")
	}
	user.ID = id
	return c.JSON(user)
}

// @Summary Delete user by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("Invalid user ID")
	}
	return c.SendStatus(204)
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
