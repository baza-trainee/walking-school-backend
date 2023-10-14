package handler

import (
	"context"
	"log/slog"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type FeedbackServiceInterface interface {
	CreateFormService(context.Context, model.Form) error
}

// @Summary Create form.
// Description Creates form.
// @Tags form
// @Accept json
// @Produce json
// @Param Form body model.CreateFormSwagger true "Form"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /form [post].
func CreateFormHandler(s FeedbackServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		form := model.Form{}

		if err := c.BodyParser(&form); err != nil {
			log.Debug("CreateFormHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(form); err != nil {
			log.Debug("CreateFormHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.CreateFormService(c.UserContext(), form); err != nil {
			return handleError(log, "CreateFormService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}
