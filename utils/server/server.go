package server

import (
	"log"

	//"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServer(a *fiber.App) {

	// Run server.
	// if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
	if err := a.Listen(":9090"); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
