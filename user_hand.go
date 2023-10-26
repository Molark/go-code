package main

import (
	"github.com/gofiber/fiber/v2"
)

type UserProfile struct {
	ID        int
	Username  string
	Firstname string
	LastName  string
	Email     string
}

func userRead(c *fiber.Ctx) error {
	user := UserProfile{
		ID:        1,
		Username:  "Molark",
		Firstname: "Egor",
		LastName:  "Botalov",
		Email:     "botalov2006egor@gmail.com",
	}

	return c.JSON(user)
}
