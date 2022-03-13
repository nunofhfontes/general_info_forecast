package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJwt() {

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	jwt, err := token.SignedString([]byte("AllYourBase")) // privateKey
	println("generated token, which is a pointer -> ", token)
	println("generated t -> ", jwt)
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		//return c.SendStatus(fiber.StatusInternalServerError)
	}

	// return c.JSON(fiber.Map{"token": t})

}
