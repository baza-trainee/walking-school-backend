package middleware

import (
	"net/http"

	"github.com/baza-trainee/walking-school-backend/internal/api/handler"
	"github.com/baza-trainee/walking-school-backend/internal/config"
	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/baza-trainee/walking-school-backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

const (
	authorization = "Authorization"
)

func NewIdentity(cfg config.AuthConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies(model.AccessCookieName)
		if token == "" {
			return c.SendStatus(http.StatusUnauthorized)
		}

		claim, err := service.ParseToken(token, cfg.SigningKey)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		c.Locals(handler.AdminID, claim.ID)
		c.Locals("refreshString", token)

		return c.Next()
	}
}
