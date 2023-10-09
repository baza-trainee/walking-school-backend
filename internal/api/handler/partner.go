package handler

import (
	"context"
	"log/slog"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type PartnerServiceInterface interface {
	CreatePartnerService(context.Context, model.Partner) error
	GetAllPartnerService(context.Context, model.PartnerQuery) ([]model.Partner, error)
	GetPartnerByIDService(context.Context, string) (model.Partner, error)
	UpdatePartnerByIDService(context.Context, model.Partner) error
	DeletePartnerByIDService(context.Context, string) error
}

// @Summary Create partner .
// Description Creates partner.
// @Tags partner
// @Accept json
// @Produce json
// @Param Partner body model.CreatePartnerSwagger true "Partner"
// @Success 201 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /partner [post].
func CreatePartnerHandler(s PartnerServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		partner := model.Partner{}

		if err := c.BodyParser(&partner); err != nil {
			log.Debug("CreatePartnerHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(partner); err != nil {
			log.Debug("CreatePartnerHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.CreatePartnerService(c.UserContext(), partner); err != nil {
			return handleError(log, "CreatePartnerService error: ", err)
		}

		return c.Status(fiber.StatusCreated).JSON(model.SetResponse(fiber.StatusCreated, "created"))
	}
}

// @Summary Get all partners.
// Description Get all partners.
// @Tags partner
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} model.Response
// @Success 204 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /partner [get].
func GetAllPartnerHandler(s PartnerServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := model.PartnerQuery{
			Limit:  standartLimitValue,
			Offset: standartOffsetValue,
		}

		if err := c.QueryParser(&query); err != nil {
			log.Debug("GetAllPartnerHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(query); err != nil {
			log.Debug("GetAllPartnerHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		result, err := s.GetAllPartnerService(c.UserContext(), query)
		if err != nil {
			return handleError(log, "GetAllPartnerService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// @Summary Get partner by id.
// Description Gets partner by id.
// @Tags partner
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /partner/{id} [get].
func GetPartnerByIDHandler(s PartnerServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := struct {
			ID string `params:"id" validate:"required,uuid"`
		}{}

		if err := c.ParamsParser(&param); err != nil {
			log.Debug("GetPartnerByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(param); err != nil {
			log.Debug("GetPartnerByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		result, err := s.GetPartnerByIDService(c.UserContext(), param.ID)
		if err != nil {
			return handleError(log, "GetPartnerByIDService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// @Summary Update partner by id.
// Description Updates partner by id.
// @Tags partner
// @Accept json
// @Produce json
// @Param Partner body model.UpdatePartnerSwagger true "Partner"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /partner [put].
func UpdatePartnerByIDHandler(s PartnerServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		partner := model.Partner{}

		if err := c.BodyParser(&partner); err != nil {
			log.Debug("UpdatePartnerByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(partner); err != nil {
			log.Debug("UpdatePartnerByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.UpdatePartnerByIDService(c.UserContext(), partner); err != nil {
			return handleError(log, "UpdatePartnerByIDService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}

// @Summary Delete partner by id.
// Description Deletes partner by id.
// @Tags partner
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /partner/{id} [delete].
func DeletePartnerByIDHandler(s PartnerServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := struct {
			ID string `params:"id" validate:"required,uuid"`
		}{}

		if err := c.ParamsParser(&param); err != nil {
			log.Debug("DeletePartnerByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(param); err != nil {
			log.Debug("DeletePartnerByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.DeletePartnerByIDService(c.UserContext(), param.ID); err != nil {
			return handleError(log, "DeletePartnerByIDService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}
