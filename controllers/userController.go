package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	fmt.Println("Should do Login")

	// For testing purposes- testing centralized Error Handling
	// 503 On vacation!
	//return fiber.NewError(fiber.StatusServiceUnavailable, "On vacation!")

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":       false,
		"msg":         nil,
		"info":        "general info",
		"anotherInfo": "another info",
	})
}

func DoLogout() {
	fmt.Println("Should do Logout")
}
