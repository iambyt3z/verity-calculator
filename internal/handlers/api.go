package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/iambyt3z/verity-calculator/api"
	"github.com/iambyt3z/verity-calculator/internal/middleware"
)

func Handler(app *fiber.App) {
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(middleware.Authorization)

	app.Post("/solve-verity", func(c *fiber.Ctx) error {
		var params = api.SolveVerityRequestBody{}

		err_parse := c.BodyParser(&params)

		if err_parse != nil {
			log.Printf("Error parsing JSON: %v", err_parse)

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}

		err_validate := params.Validate()

		if err_validate != nil {
			log.Printf("Error parsing JSON: %v", err_parse)

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err_validate.Error(),
			})
		}

		var response = api.SolveVerityResponse{
			OutsideDissectionSteps: []string{},
			InsideDissectionSteps:  [][]string{},
		}

		return c.Status(fiber.StatusOK).JSON(response)
	})
}
