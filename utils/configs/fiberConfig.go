package fiberConfig

import (
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

			err = ctx.Status(code).SendString("Experimental Error String.")
			if err != nil {
				return err
			}

			// Return from handler
			return nil
		},
	}
}
