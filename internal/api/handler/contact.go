package handler

import (
	"context"
	"log/slog"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type ContactServiceInterface interface {
	CreateContactService(context.Context, model.Contact) error
	GetAllContactService(context.Context) ([]model.Contact, error)
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

// @Summary Get all contacts.
// Description Get all contacts.
// @Tags contact
// @Produce json
// @Success 200 {object} model.Response
// @Success 204 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /contact [get].
func GetAllContactHandler(s ContactServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := s.GetAllContactService(c.UserContext())
		if err != nil {
			return handleError(log, "GetAllContactService error: ", err)
		}

		if len(result) < minimalResult {
			return c.Status(fiber.StatusNoContent).JSON(model.SetResponse(fiber.StatusNoContent, "no content"))
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
