package handler

import (
	"context"
	"log/slog"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type ImagesCarouselServiceInterface interface {
	CreateImagesCarouselService(context.Context, model.ImageCarousel) error
	GetAllImagesCarouselService(context.Context, model.ImageCarouselQuery) ([]model.ImageCarousel, error)
	GetImagesCarouselByIDService(context.Context, string) (model.ImageCarousel, error)
	DeleteImagesCarouselByIDService(context.Context, string) error
}

// @Summary Create image.
// Description Creates image.
// @Tags image carousel
// @Accept json
// @Produce json
// @Param ImagesCarousel body model.CreateImageCarouselSwagger true "ImagesCarousel"
// @Success 201 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /image-carousel [post].
func CreateImagesCarouselHandler(s ImagesCarouselServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		imagesCarousel := model.ImageCarousel{}

		if err := c.BodyParser(&imagesCarousel); err != nil {
			log.Debug("CreateImagesCarouselHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(imagesCarousel); err != nil {
			log.Debug("CreateImagesCarouselHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.CreateImagesCarouselService(c.UserContext(), imagesCarousel); err != nil {
			return handleError(log, "CreateImagesCarouselService error: ", err)
		}

		return c.Status(fiber.StatusCreated).JSON(model.SetResponse(fiber.StatusCreated, "created"))
	}
}

// @Summary Get all images.
// Description Get all images.
// @Tags image carousel
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} model.Response
// @Success 204 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /image-carousel [get].
func GetAllImagesCarouselHandler(s ImagesCarouselServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := model.ImageCarouselQuery{
			Limit:  standartLimitValue,
			Offset: standartOffsetValue,
		}

		if err := c.QueryParser(&query); err != nil {
			log.Debug("GetAllImagesCarouselHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(query); err != nil {
			log.Debug("GetAllImagesCarouselHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		result, err := s.GetAllImagesCarouselService(c.UserContext(), query)
		if err != nil {
			return handleError(log, "GetAllImagesCarouselService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// @Summary Get image by id.
// Description Gets image by id.
// @Tags image carousel
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /image-carousel/{id} [get].
func GetImagesCarouselByIDHandler(s ImagesCarouselServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := struct {
			ID string `params:"id" validate:"required,uuid"`
		}{}

		if err := c.ParamsParser(&param); err != nil {
			log.Debug("GetImagesCarouselByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(param); err != nil {
			log.Debug("GetImagesCarouselByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		result, err := s.GetImagesCarouselByIDService(c.UserContext(), param.ID)
		if err != nil {
			return handleError(log, "GetImagesCarouselByIDService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// @Summary Delete image by id.
// Description Deletes image by id.
// @Tags image carousel
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /image-carousel/{id} [delete].
func DeleteImagesCarouselByIDHandler(s ImagesCarouselServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := struct {
			ID string `params:"id" validate:"required,uuid"`
		}{}

		if err := c.ParamsParser(&param); err != nil {
			log.Debug("DeleteImagesCarouselByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(param); err != nil {
			log.Debug("DeleteImagesCarouselByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.DeleteImagesCarouselByIDService(c.UserContext(), param.ID); err != nil {
			return handleError(log, "DeleteImagesCarouselByIDService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}
