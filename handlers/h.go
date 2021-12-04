package handlers

import (
	"server/validate"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	t, err := validate.GetAccessToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.ErrInternalServerError)
	}
	return c.JSON(fiber.Map{
		"token": t,
	})
}

func Profile(c *fiber.Ctx) error {
	user := c.Locals("aud")
	return c.JSON(fiber.Map{
		"info": "profile",
		"user": user,
	})
}
