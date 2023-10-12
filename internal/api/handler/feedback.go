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
// @Param Feedback body model.CreateFeedbackSwagger true "Feedback"
// @Success 201 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /feedback [post].
func CreateFeedbackHandler(s FeedbackServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		feedback := model.Feedback{}

		if err := c.BodyParser(&feedback); err != nil {
			log.Debug("CreateFeedbackHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(feedback); err != nil {
			log.Debug("CreateFeedbackHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.CreateFeedbackService(c.UserContext(), feedback); err != nil {
			return handleError(log, "CreateFeedbackService error: ", err)
		}

		return c.Status(fiber.StatusCreated).JSON(model.SetResponse(fiber.StatusCreated, "created"))
	}
}
