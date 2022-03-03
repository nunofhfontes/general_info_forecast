package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouter(app *fiber.App) {

	// Register the Routes here

	PublicRoutes(app)
}

type RouterRegistry struct {
}

func (s *RouterRegistry) RegisterPublicRoutes(str string) *RouterRegistry {

	fmt.Println("Registering Public Routes, str: ", str)

	return s
}

func (s *RouterRegistry) RegisterPrivateRoutes(str string) *RouterRegistry {

	fmt.Println("Registering Private Routes, str: ", str)

	return s
}

func (s *RouterRegistry) RegisterNotFoundRoute(str string) *RouterRegistry {

	fmt.Println("Registering Not Found Route, str: ", str)

	return s
}
