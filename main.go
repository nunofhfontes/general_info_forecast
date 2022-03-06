package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	//fUtils "dataForecast/utils/files"
	"fmt"

	middleware "dataForecast/middleware"
	router "dataForecast/utils/router"
	serverUtils "dataForecast/utils/server"

	fiberConfig "dataForecast/utils/configs"
)

func main() {
	fmt.Println("Initializing REST API...")

	// Define Fiber config.
	config := fiberConfig.ConfigFiber()

	// New Fiber
	app := fiber.New(config)

	// recovers from panics anywhere in the stack chain and handles the control to the centralized ErrorHandler.
	app.Use(recover.New())

	// Register Middleware
	middleware.AppMiddleware(app)

	// Register Router
	// router.RegisterRouter(app)

	// testing purposes method chaining
	routerRegister := router.RouterRegistry{}
	routerRegister.
		RegisterPublicRoutes(app, "public").
		RegisterPrivateRoutes("private").
		RegisterNotFoundRoute("NotFOund")

	// Start Server
	serverUtils.StartServer(app)
}
