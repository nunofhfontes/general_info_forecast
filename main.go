package main

import (
	quoteService "dataForecast/services"
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

	quoteService.GetQuote()

	newConfig, err := fiberConfig.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	fmt.Printf("new config : %s = %s \n", "Migration URL", newConfig.MigrationURL)

	db.InitCockroachDB(newConfig)

	//get config file and config application
	fiberConfig.GetAndConfigApp()

	fmt.Println("Initializing REST API...")

	// Initialize DB
	//db.InitDB()

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
