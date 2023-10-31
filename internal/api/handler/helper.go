package handler

import (
	"errors"
	"log/slog"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New() //nolint

const (
	standartLimitValue  = 10
	standartOffsetValue = 0
	minimalResult       = 1
)

func UserValidate(validate *validator.Validate, user model.User) error {
	if user.Phone == "" && user.Email == "" {
		return errors.New("phone or email must be input")
	}

	if err := validate.Struct(user); err != nil {
		return err
	}

	return nil
}

func handleError(log *slog.Logger, message string, err error) error {
	switch {
	case errors.Is(err, model.ErrRequestTimeout):
		log.Debug(message, err.Error())

		return fiber.NewError(fiber.StatusRequestTimeout, err.Error())
	case errors.Is(err, model.ErrConflict):
		log.Debug(message, err.Error())

		return fiber.NewError(fiber.StatusConflict, err.Error())
	case errors.Is(err, model.ErrNotFound):
		log.Debug(message, err.Error())

		return fiber.NewError(fiber.StatusNotFound, err.Error())
	default:
		log.Error(message, err.Error())

		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
}
