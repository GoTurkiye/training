package main

import (
	"training/egitim_kampi"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(egitim_kampi.Public)
	})

	app.Listen(":3000")
}
