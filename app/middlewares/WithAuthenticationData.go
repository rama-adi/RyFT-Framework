package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rama-adi/RyFT-Framework/app/models"
	"github.com/rama-adi/RyFT-Framework/app/utils"
	"strings"
)

func WithAuthenticationData(c *fiber.Ctx) error {

	authorizationHeader := c.Get("Authorization", "")

	if authorizationHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.HttpResponse{
			Success: false,
			Message: "Missing Authorization header",
			Data:    nil,
		})
	}

	rep := strings.Replace(authorizationHeader, "Bearer ", "", 1)
	user, err := models.User{}.FromAccessToken(rep)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.HttpResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	c.Locals("accessToken", rep)
	c.Locals("user", user)

	return c.Next()
}
