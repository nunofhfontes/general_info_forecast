package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
)

func AppMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route
		cors.New(),
		// Add logger
		logger.New(),

		// Helmet middleware helps to secure our Fiber application by setting various HTTP headers
		helmet.New(), // add Helmet middleware

		// FIXME - testing new logger format
		// logger.New(logger.Config{
		// 	Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		// }),
	)
}
