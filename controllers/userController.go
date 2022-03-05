package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	fmt.Println("Should do Login")

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
