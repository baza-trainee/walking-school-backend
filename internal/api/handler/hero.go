package handler

import (
	"context"
	"log/slog"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type HeroServiceInterface interface {
	CreateHeroService(context.Context, model.Hero) error
	GetAllHeroService(context.Context, model.HeroQuery) ([]model.Hero, error)
	GetHeroByIDService(context.Context, string) (model.Hero, error)
	UpdateHeroByIDService(context.Context, model.Hero) error
	DeleteHeroByIDService(context.Context, string) error
}

// @Summary Create hero .
// Description Creates hero.
// @Tags hero
// @Accept json
// @Produce json
// @Param Hero body model.CreateHeroSwagger true "Hero"
// @Success 201 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /hero [post].
func CreateHeroHandler(s HeroServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		hero := model.Hero{}

		if err := c.BodyParser(&hero); err != nil {
			log.Debug("CreateHeroHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(hero); err != nil {
			log.Debug("CreateHeroHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.CreateHeroService(c.UserContext(), hero); err != nil {
			return handleError(log, "CreateHeroService error: ", err)
		}

		return c.Status(fiber.StatusCreated).JSON(model.SetResponse(fiber.StatusCreated, "created"))
	}
}

// @Summary Get all heros.
// Description Get all heros.
// @Tags hero
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} model.Response
// @Success 204 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /hero [get].
func GetAllHeroHandler(s HeroServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := model.HeroQuery{
			Limit:  standartLimitValue,
			Offset: standartOffsetValue,
		}

		if err := c.QueryParser(&query); err != nil {
			log.Debug("GetAllHeroHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(query); err != nil {
			log.Debug("GetAllHeroHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		result, err := s.GetAllHeroService(c.UserContext(), query)
		if err != nil {
			return handleError(log, "GetAllHeroService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// @Summary Get hero by id.
// Description Gets hero by id.
// @Tags hero
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /hero/{id} [get].
func GetHeroByIDHandler(s HeroServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := struct {
			ID string `params:"id" validate:"required,uuid"`
		}{}

		if err := c.ParamsParser(&param); err != nil {
			log.Debug("GetHeroByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(param); err != nil {
			log.Debug("GetHeroByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		result, err := s.GetHeroByIDService(c.UserContext(), param.ID)
		if err != nil {
			return handleError(log, "GetHeroByIDService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// @Summary Update hero by id.
// Description Updates hero by id.
// @Tags hero
// @Accept json
// @Produce json
// @Param Hero body model.UpdateHeroSwagger true "Hero"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /hero [put].
func UpdateHeroByIDHandler(s HeroServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		hero := model.Hero{}

		if err := c.BodyParser(&hero); err != nil {
			log.Debug("UpdateHeroByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(hero); err != nil {
			log.Debug("UpdateHeroByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.UpdateHeroByIDService(c.UserContext(), hero); err != nil {
			return handleError(log, "UpdateHeroByIDService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}

// @Summary Delete hero by id.
// Description Deletes hero by id.
// @Tags hero
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /hero/{id} [delete].
func DeleteHeroByIDHandler(s HeroServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := struct {
			ID string `params:"id" validate:"required,uuid"`
		}{}

		if err := c.ParamsParser(&param); err != nil {
			log.Debug("DeleteHeroByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(param); err != nil {
			log.Debug("DeleteHeroByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.DeleteHeroByIDService(c.UserContext(), param.ID); err != nil {
			return handleError(log, "DeleteHeroByIDService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}
