package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	datamodels "dataForecast/models"
	authUtils "dataForecast/utils/auth"
)

func Login(c *fiber.Ctx) error {
	fmt.Println("Should do Login")

	// For testing purposes- testing centralized Error Handling
	// 503 On vacation!
	//return fiber.NewError(fiber.StatusServiceUnavailable, "Testing Error MessageXYZ!")

	// user := c.FormValue("user")
	// pass := c.FormValue("password")

	payload := datamodels.AuthInfo{}
	c.BodyParser(&payload)

	println("user: ", payload.User)
	println("password: ", payload.Password)

	// 1 - get data from DB

	// 2 - put it into a structure

	// get an access toke (jwt) with the submited info
	authUtils.CreateJwt()

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

func Register(c *fiber.Ctx) error {

	return fiber.NewError(fiber.StatusServiceUnavailable, "Register is not Working yet!")
}

func DoLogout() {
	fmt.Println("Should do Logout")
}
