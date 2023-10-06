package handler

import (
	"context"
	"errors"
	"log/slog"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type UserServiceInterface interface {
	CreateUserService(context.Context, model.User) error
	GetAllUserService(context.Context, model.UserQuery) ([]model.User, error)
}

// @Summary Create user.
// Description Creates user.
// @Tags user
// @Accept json
// @Produce json
// @Param User body model.CreateUserSwagger true "User"
// @Success 200 {object} model.Response
// @Success 201 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 409 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /user [post].
func CreateUserHandler(s UserServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := model.User{}

		if err := c.BodyParser(&user); err != nil {
			log.Debug("CreateUserHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := UserValidate(validate, user); err != nil {
			log.Debug("CreateUserHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.CreateUserService(c.UserContext(), user); err != nil {
			if errors.Is(err, model.ErrConflict) {
				log.Debug("CreateUserService error: ", err.Error())

				return fiber.NewError(fiber.StatusConflict, err.Error())
			}

			log.Error("CreateUserService error: ", err.Error())
			// Какие ошибки могут возвращаться?
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(model.SetResponse(fiber.StatusCreated, "created"))
	}
}

// @Summary Get all users.
// Description Get all users.
// @Tags user
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} model.Response
// @Success 204 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /user [get].
func GetAllUserHandler(s UserServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := model.UserQuery{
			Limit:  standartLimitValue,
			Offset: standartOffsetValue,
		}

		if err := c.QueryParser(&query); err != nil {
			log.Debug("GetAllUserHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(query); err != nil {
			log.Debug("GetAllUserHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		result, err := s.GetAllUserService(c.UserContext(), query)
		if err != nil {
			if errors.Is(err, model.ErrNoContent) {
				log.Debug("GetAllUserService error: ", err.Error())

				return fiber.NewError(fiber.StatusNoContent, err.Error())
			}

			log.Error("GetAllUserService error: ", err.Error())
			// Какие ошибки могут возвращаться?
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}
