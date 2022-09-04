package routes

import (
	"github.com/akshay1027/url-shortner/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	//All routes related to urls comes here
	app.Get("api/v1/:url", controllers.ResolveURL)
	app.Post("/api/v1", controllers.ShortenURL)
}
