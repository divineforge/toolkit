package divineforge/toolkit

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/swaggo/swag/cmd/swag"
)

func main() {
	app := fiber.New()

	// Swagger setup
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	// ...
}
