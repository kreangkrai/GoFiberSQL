package Router

import (
	"github.com/gofiber/fiber/v2"
)

func DeviceMiddleware() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		c.SendString("Device")
	}
}
func LocationMiddleware() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		c.SendString("Location")
	}
}
