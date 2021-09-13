package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	log.Println("[*] Starting Onlyballs Server...")

	// Create the Fiber worker so we can add routes to it
	app := fiber.New()

	// Set up logging functionality
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(favicon.New())

	// Or extend your config for customization
	app.Use(favicon.New(favicon.Config{
		File: "./public/assets/img/favicon.ico",
	}))

	app.Static("/", "./public")

	app.Post("/login", LoginHandler)

	// Run the service
	log.Fatalln(app.Listen(":8008"))

	log.Println("[*] Shutting Onlyballs Server down...")
}

func IndexHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func LoginHandler(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
