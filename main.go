package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/iambyt3z/verity-calculator/internal/handlers"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("error loading .env file:", err)
	} else {
		fmt.Println(".env file loaded successfully")
	}
}

func main() {
	loadEnv()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, https://d2-verity-calculator.netlify.app/",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	handlers.Handler(app)

	fmt.Println("Starting Verity Calculator API Service....")

	app.Listen(":3000")
}
