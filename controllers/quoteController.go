package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetQuote(c *fiber.Ctx) error {

	fmt.Println("Should get a Quote")

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":       false,
		"msg":         nil,
		"info":        "general info",
		"anotherInfo": "another info",
	})
}
