package router

import (
	quoteController "dataForecast/controllers"
	middleware "dataForecast/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(app *fiber.App) {

	// Create routes group.
	route := app.Group("/api/v1")

	route.Get("/quote", middleware.JWTProtected(), quoteController.GetQuote)
	// route.Get("/quote", quoteController.GetQuote)

	// route.Get("/login", controllers.Login) // get list of all books

	// Routes for GET method:
	//route.Get("/home", controllers.GetBooks) // get list of all books
	// route.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World 👋!")
	// })

	// app.Get("/", func(c *fiber.Ctx) error {
	//     return c.SendString("Hello, World 👋!")
	// })
}
