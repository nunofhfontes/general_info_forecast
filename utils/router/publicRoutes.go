package router

import (
	"dataForecast/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	route.Get("/login", controllers.Login)

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
