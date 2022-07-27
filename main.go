package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"log"
	"os"

	//fUtils "dataForecast/utils/files"
	"fmt"

	middleware "dataForecast/middleware"
	router "dataForecast/utils/router"
	serverUtils "dataForecast/utils/server"

	db "dataForecast/db"
	fiberConfig "dataForecast/utils/configs"
)

func main() {

	// Load env variables
	// godotenv package
	dotenv := goDotEnvVariable("BD_URL")
	fmt.Printf("godotenv : %s = %s \n", "BD_URL", os.Getenv("BD_URL"))
	fmt.Printf("godotenv os.getenv: %s = %s \n", "BD_URL", dotenv)

	fmt.Println("Initializing REST API...")

	// Initialize DB
	db.InitDB()

	// Define Fiber config.
	config := fiberConfig.ConfigFiber()

	// New Fiber
	app := fiber.New(config)

	app.Static(
		"/static",  // mount address
		"./public", // path to the file folder
	)

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
		RegisterPrivateRoutes(app, "private").
		RegisterNotFoundRoute("NotFOund")

	// Start Server
	serverUtils.StartServer(app)
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
