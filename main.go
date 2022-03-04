package main

import (
	"github.com/gofiber/fiber/v2"
	//fUtils "dataForecast/utils/files"
	"fmt"

	middleware "dataForecast/middleware"
	router "dataForecast/utils/router"
	serverUtils "dataForecast/utils/server"
)

func main() {
	fmt.Println("Initializing REST API...")

	// New Fiber
	app := fiber.New()

	// Register Middleware
	middleware.AppMiddleware(app)

	// Register Router
	router.RegisterRouter(app)

	// testing purposes method chaining
	routerRegister := router.RouterRegistry{}
	routerRegister.
		RegisterPublicRoutes("public").
		RegisterPrivateRoutes("private").
		RegisterNotFoundRoute("NotFOund")

	// Start Server
	serverUtils.StartServer(app)
}
