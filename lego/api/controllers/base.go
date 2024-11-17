package controllers

import (
	"github.com/gofiber/fiber/v2"
	C "other-side/config"
)

func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"ok":  1,
		"v":   C.Conf.Version,
		"env": C.Conf.Environment,
	})
}
