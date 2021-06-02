package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vivekv96/go-admin/utils"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := utils.ParseJWT(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated user",
		})
	}

	return c.Next()
}
