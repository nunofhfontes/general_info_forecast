package router

import (
	userController "dataForecast/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	// Create routes group.
	route := app.Group("/api/v1")

	//FIXME - leaving this route just for testing purposes
	route.Get("/login", userController.Login)

	route.Post("/login", userController.Login)

	route.Get("/register", userController.Register)

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
