package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akshay1027/url-shortner/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	//app.Use(csrf.New())
	app.Use(logger.New())

	routes.UserRoute(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Listen to port
	log.Fatal(app.Listen(os.Getenv("PORT")))
}
