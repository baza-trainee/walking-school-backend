package handler

import (
	"context"
	"log/slog"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type ContactServiceInterface interface {
	CreateContactService(context.Context, model.Contact) error
	GetContactByIDService(context.Context, string) (model.Contact, error)
	UpdateContactByIDService(context.Context, model.Contact) error
}

// @Summary Create contact .
// Description Creates contact.
// @Tags contact
// @Accept json
// @Produce json
// @Param Contact body model.CreateContactSwagger true "Contact"
// @Success 201 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /contact [post].
func CreateContactHandler(s ContactServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		contact := model.Contact{}

		if err := c.BodyParser(&contact); err != nil {
			log.Debug("CreateContactHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(contact); err != nil {
			log.Debug("CreateContactHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.CreateContactService(c.UserContext(), contact); err != nil {
			return handleError(log, "CreateContactService error: ", err)
		}

		return c.Status(fiber.StatusCreated).JSON(model.SetResponse(fiber.StatusCreated, "created"))
	}
}

// @Summary Get contact by id.
// Description Gets contact by id.
// @Tags contact
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /contact/{id} [get].
func GetContactByIDHandler(s ContactServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := struct {
			ID string `params:"id" validate:"required,uuid"`
		}{}

		if err := c.ParamsParser(&param); err != nil {
			log.Debug("GetContactByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(param); err != nil {
			log.Debug("GetContactByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		result, err := s.GetContactByIDService(c.UserContext(), param.ID)
		if err != nil {
			return handleError(log, "GetContactByIDService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// @Summary Update contact by id.
// Description Updates contact by id.
// @Tags contact
// @Accept json
// @Produce json
// @Param Contact body model.UpdateContactSwagger true "Contact"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /contact [put].
func UpdateContactByIDHandler(s ContactServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		contact := model.Contact{}

		if err := c.BodyParser(&contact); err != nil {
			log.Debug("UpdateContactByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(contact); err != nil {
			log.Debug("UpdateContactByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.UpdateContactByIDService(c.UserContext(), contact); err != nil {
			return handleError(log, "UpdateContactByIDService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}
