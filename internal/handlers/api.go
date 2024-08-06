package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/iambyt3z/verity-calculator/internal/middleware"
)

func Handler(app *fiber.App) {
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(middleware.Authorization)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello World!",
		})
	})
}
