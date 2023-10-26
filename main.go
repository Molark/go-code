package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("value"))

	})
	app.Listen(":8080")

}
