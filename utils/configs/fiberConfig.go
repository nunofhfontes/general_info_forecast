package fiberConfig

import (
	datamodels "dataForecast/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ConfigFiber() fiber.Config {

	return fiber.Config{

		ServerHeader: "REST API - Go -> Fiber",
		AppName:      "Test App v1.00",

		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			// TODO - Parse all types of Errors that can be thrown on app's layers

			// Set Content-Type: application/json; charset=utf-8
			ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

			// Returning an entire JSON struct, ErrorResponse, with error messaged, code, etc etc
			err = ctx.Status(code).JSON(datamodels.ErrorResponse{strconv.Itoa(code), err.Error()})
			// err = ctx.Status(code).SendString(err.Error())
			if err != nil {
				return err
			}

			// Return from handler
			return nil
		},
	}
}
