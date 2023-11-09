package handler

import (
	"context"
	"log/slog"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type FeedbackServiceInterface interface {
	CreateFeedbackService(context.Context, model.Feedback) error
}

// @Summary Create feedback.
// Description Creates feedback.
// @Tags feedback
// @Accept json
// @Produce json
// @Param feedback body model.CreateFeedbackSwagger true "feedback"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /feedback [post].
func CreateFeedbackHandler(s FeedbackServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		form := model.Feedback{}

		if err := c.BodyParser(&form); err != nil {
			log.Debug("CreateFeedbackHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(form); err != nil {
			log.Debug("CreateFeedbackHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.CreateFeedbackService(c.UserContext(), form); err != nil {
			return handleError(log, "CreateFeedbackService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}
