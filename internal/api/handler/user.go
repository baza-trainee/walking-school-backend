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
	GetUserByIDService(context.Context, string) (model.User, error)
	UpdateUserByIDService(context.Context, model.User) error
	DeleteUserByIDService(context.Context, string) error
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

// @Summary Get user by id.
// Description Gets user by id.
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /user/{id} [get].
func GetUserByIDHandler(s UserServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := struct {
			ID string `params:"id" validate:"required,uuid"`
		}{}

		if err := c.ParamsParser(&param); err != nil {
			log.Debug("GetUserByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(param); err != nil {
			log.Debug("GetUserByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		result, err := s.GetUserByIDService(c.UserContext(), param.ID)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				log.Debug("GetUserByIDService error: ", err.Error())

				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}

			log.Error("GetUserByIDService error: ", err.Error())
			// Какие ошибки могут возвращаться?
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// @Summary Update user by id.
// Description Updates user by id.
// @Tags user
// @Accept json
// @Produce json
// @Param User body model.UpdateUserSwagger true "User"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /user [put].
func UpdateUserByIDHandler(s UserServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := model.User{}

		if err := c.BodyParser(&user); err != nil {
			log.Debug("UpdateUserByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := UserValidate(validate, user); err != nil {
			log.Debug("CreateUserHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.UpdateUserByIDService(c.UserContext(), user); err != nil {
			if errors.Is(err, model.ErrConflict) {
				log.Debug("UpdateUserByIDService error: ", err.Error())

				return fiber.NewError(fiber.StatusConflict, err.Error())
			}

			if errors.Is(err, model.ErrNotFound) {
				log.Debug("UpdateUserByIDService error: ", err.Error())

				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}

			log.Error("UpdateUserByIDService error: ", err.Error())
			// Какие ошибки могут возвращаться?
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}

// @Summary Delete user by id.
// Description Deletes user by id.
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /user/{id} [delete].
func DeleteUserByIDHandler(s UserServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := struct {
			ID string `params:"id" validate:"required,uuid"`
		}{}

		if err := c.ParamsParser(&param); err != nil {
			log.Debug("DeleteUserByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(param); err != nil {
			log.Debug("DeleteUserByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.DeleteUserByIDService(c.UserContext(), param.ID); err != nil {
			if errors.Is(err, model.ErrNotFound) {
				log.Debug("DeleteUserByIDService error: ", err.Error())

				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}

			log.Error("DeleteUserByIDService error: ", err.Error())
			// Какие ошибки могут возвращаться?
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}
