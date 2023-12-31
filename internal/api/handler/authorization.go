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
	ForgotPasswordService(context.Context, string) error
	ResetPasswordService(context.Context, model.ResetPassword) error
	RegistrationForTestService(context.Context, model.Admin) error
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

// @Summary Forgot password.
// @Description Reset password.
// @Tags authorization
// @Accept json
// @Produce json
// @Param Email body model.Login true "Email to authenticate user"
// @Success 200 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /forgot-password [post].
func ForgotPasswordHandler(s AuthorizationServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		login := model.Login{}

		if err := c.BodyParser(&login); err != nil {
			log.Debug("ForgotPasswordHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(login); err != nil {
			log.Debug("ForgotPasswordHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.ForgotPasswordService(c.UserContext(), login.Login); err != nil {
			return handleError(log, "ForgotPasswordService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}

// @Summary Reset password.
// @Description Reset password.
// @Tags authorization
// @Accept json
// @Produce json
// @Param ResetPassword body model.ResetPassword true "Reset password to access to account"
// @Success 200 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /reset-password [post].
func ResetPasswordHandler(s AuthorizationServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := model.ResetPassword{}

		if err := c.BodyParser(&data); err != nil {
			log.Debug("ResetPasswordHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(data); err != nil {
			log.Debug("ResetPasswordHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if data.NewPassword != data.ConfirmedNewPassword {
			log.Debug("ResetPasswordHandler error: ", "new password and confirmed one don't match")

			return fiber.NewError(fiber.StatusBadRequest, "new password and confirmed one don't match")
		}

		if err := s.ResetPasswordService(c.UserContext(), data); err != nil {
			return handleError(log, "ResetPasswordService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}

// @Summary Registration for test.
// @Description Registration for test.
// @Tags authorization
// @Accept json
// @Produce json
// @Param RegistrationForTest body model.Identity true "Registration for test"
// @Success 201 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /registration-for-test [post].
func RegistrationForTestHandler(s AuthorizationServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		admin := model.Admin{}

		if err := c.BodyParser(&admin); err != nil {
			log.Debug("RegistrationForTestHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(admin); err != nil {
			log.Debug("RegistrationForTestHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.RegistrationForTestService(c.UserContext(), admin); err != nil {
			return handleError(log, "RegistrationForTestService error: ", err)
		}

		return c.Status(fiber.StatusCreated).JSON(model.SetResponse(fiber.StatusCreated, "created"))
	}
}
