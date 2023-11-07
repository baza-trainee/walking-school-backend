package handler

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/baza-trainee/walking-school-backend/internal/config"
	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type AuthorizationServiceInterface interface {
	SignInService(context.Context, model.Identity) (model.TokenPair, error)
	RefreshService(context.Context, string) (model.TokenPair, error)
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
func SignInHandler(s AuthorizationServiceInterface, log *slog.Logger, cfg config.AuthConfig) fiber.Handler {
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

				return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
			}

			return handleError(log, "SignInService error: ", err)
		}

		c.Cookie(newCookie(
			model.AccessCookieName,
			result.AccessToken,
			model.AccessCookiePath,
			time.Now().Add(cfg.AccessTokenTTL),
		))

		c.Cookie(newCookie(
			model.RefreshCookieName,
			result.RefreshToken,
			model.RefreshCookiePath,
			time.Now().Add(cfg.RefreshTokenTTL),
		))

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}

// @Summary Logout.
// @Description Cleaning of access and refresh cookies for authorized user.
// @Tags authorization
// @Produce json
// @Success 200 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /logout [post].
func SignOutHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Cookies(model.AccessCookieName) == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Cookie(newCookie(
			model.AccessCookieName,
			"",
			model.AccessCookiePath,
			time.Now().Add(-1),
		))

		c.Cookie(newCookie(
			model.RefreshCookieName,
			"",
			model.RefreshCookiePath,
			time.Now().Add(-1),
		))

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}

// @Summary Refreshing tokens.
// @Description Renew accept and refresh tokens.
// @Tags authorization
// @Accept json
// @Produce json
// @Param TokenPair body model.TokenPair true "couple of access and refresh tokens"
// @Success 200 {object} model.TokenPair
// @Failure 401 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /authorization-refresh [post].
func RefreshHandler(s AuthorizationServiceInterface, log *slog.Logger, cfg config.AuthConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		refreshToken := c.Cookies(model.RefreshCookieName)
		if refreshToken == "" {
			log.Debug(fmt.Sprintf("%s is empty", model.RefreshCookieName))

			return c.SendStatus(fiber.StatusUnauthorized)
			// return fiber.NewError(fiber.StatusUnauthorized, fmt.Sprintf("%s is empty", model.RefreshCookieName))
		}

		result, err := s.RefreshService(c.UserContext(), refreshToken)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				log.Debug("RefreshService error: ", err.Error())

				return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
			}

			return handleError(log, "RefreshService error: ", err)
		}

		c.Cookie(newCookie(
			model.AccessCookieName,
			result.AccessToken,
			model.AccessCookiePath,
			time.Now().Add(cfg.AccessTokenTTL),
		))

		c.Cookie(newCookie(
			model.RefreshCookieName,
			result.RefreshToken,
			model.RefreshCookiePath,
			time.Now().Add(cfg.RefreshTokenTTL),
		))

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}
