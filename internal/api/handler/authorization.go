package handler

import (
	"context"
	"errors"
	"log/slog"
	"strings"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type AuthorizationServiceInterface interface {
	SignInService(context.Context, model.Identity) (model.TokenPair, error)
}

// @Summary Authorization.
// @Description Accepts email and password to authorize the admin.
// @Tags authorization
// @Accept json
// @Produce json
// @Param Identity body model.Identity true "email and password to login"
// @Success 200 {object} model.TokenPair
// @Failure 401 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /login [post].
func SignInHandler(s AuthorizationServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		identity := model.Identity{}

		if err := c.BodyParser(&identity); err != nil {
			log.Debug("SignInHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(identity); err != nil {
			log.Debug("SignInHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		identity.Login = strings.ToLower(identity.Login)

		result, err := s.SignInService(c.UserContext(), identity)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				log.Debug("SignInService error: ", err.Error())

				return fiber.NewError(fiber.StatusUnauthorized, err.Error())
			}

			return handleError(log, "SignInService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}
