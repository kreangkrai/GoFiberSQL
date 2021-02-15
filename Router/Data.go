package Router

import (
	"GoFiberSQL/Controller"
	"GoFiberSQL/Models"

	"github.com/gofiber/fiber/v2"
)

func Insert(c *fiber.Ctx) error {
	var input Models.Data
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	m, err := Controller.Insert(input)
	if err != nil {
		return c.JSON(err.Error())
	}

	return c.JSON(m)
}

func Get(c *fiber.Ctx) error {
	device := c.Params("device")
	data, err := Controller.GetData(device)
	if err != nil {
		return nil
	}
	return c.JSON(data)
}
func Gets(c *fiber.Ctx) error {
	data, err := Controller.GetsData()
	if err != nil {
		return nil
	}
	return c.JSON(data)
}
func Update(c *fiber.Ctx) error {
	var input Models.Data

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	m, err := Controller.Update(input)
	if err != nil {
		return c.JSON(err.Error())
	}
	return c.JSON(m)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	m, err := Controller.Delete(id)
	if err != nil {
		return c.JSON(err.Error())
	}
	return c.JSON(m)
}
