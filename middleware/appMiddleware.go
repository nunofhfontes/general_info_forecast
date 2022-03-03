package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func AppMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route
		cors.New(),
		// Add logger
		logger.New(),
	)
}
