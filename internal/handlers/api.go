package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/iambyt3z/verity-calculator/api"
	"github.com/iambyt3z/verity-calculator/internal/middleware"
	"github.com/iambyt3z/verity-calculator/internal/verity_calculator"
)

func Handler(app *fiber.App) {
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(middleware.Authorization)

	corsConfig := cors.Config{
		AllowOrigins: "http://localhost:5173, https://d2-verity-calculator.netlify.app/",
		AllowMethods: "POST",
	}

	app.Use(cors.New(corsConfig))

	app.Post("/solve-verity", func(c *fiber.Ctx) error {
		var params = api.SolveVerityRequestBody{}

		err_parse := c.BodyParser(&params)

		if err_parse != nil {
			log.Printf("Error parsing JSON: %v", err_parse)

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}

		log.Println("Request JSON parsed successfully")

		err_validate := params.Validate()

		if err_validate != nil {
			log.Printf("Error parsing JSON: %v", err_parse)

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err_validate.Error(),
			})
		}

		log.Println("Request JSON validated successfully")

		log.Println("Solving verity dissection...")

		outsideDissectionSteps, outsideTargetStatueShapeNames := verity_calculator.SolveOutsideDissection(
			params.InsideRoomLeftStatueSymbol,
			params.InsideRoomMidStatueSymbol,
			params.InsideRoomRightStatueSymbol,
			params.OutsideRoomLeftStatueSymbol,
			params.OutsideRoomMidStatueSymbol,
			params.OutsideRoomRightStatueSymbol,
			params.IsChallengePhaseTwo,
		)

		log.Println("Verity dissection solved")

		var response = api.SolveVerityResponse{
			OutsideDissectionSteps:        outsideDissectionSteps,
			OutsideTargetStatueShapeNames: outsideTargetStatueShapeNames,
			InsideDissectionSteps:         [][]string{},
		}

		log.Println("Response prepared successfully")

		return c.Status(fiber.StatusOK).JSON(response)
	})
}
