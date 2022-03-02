package main

import (
	"github.com/gofiber/fiber/v2"
	//fUtils "dataForecast/utils/files"
	"fmt"

	serverUtils "dataForecast/utils/server"
)

func main() {
	fmt.Println("Initializing REST API...")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	serverUtils.StartServer(app)
}

// TODO  - realocate for other files
// 1 - read json with locations converting to a raw variable
//fUtils.ReadLocationsRaw()

// 2 - read json with locations parsing info to struts
//fUtils.ReadLocationsParsed4Strut()

// 2 - init cron jobs
//cronJobs.InitCronJobs()
