package handler

import (
	"context"
	"errors"
	"log/slog"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type ProjSectDescServiceInterface interface {
	CreateProjSectDescService(context.Context, model.ProjSectDesc) error
	GetProjSectDescByIDService(context.Context, string) (model.ProjSectDesc, error)
	UpdateProjSectDescByIDService(context.Context, model.ProjSectDesc) error
}

// @Summary Create projects section description.
// Description Create projects section description.
// @Tags projects section description
// @Accept json
// @Produce json
// @Param ProjSectDesc body model.CreateProjSectDescSwagger true "ProjSectDesc"
// @Success 201 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /projects-section-description [post].
func CreateProjSectDescHandler(s ProjSectDescServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		projSectDesc := model.ProjSectDesc{}

		if err := c.BodyParser(&projSectDesc); err != nil {
			log.Debug("CreateProjSectDescHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(projSectDesc); err != nil {
			log.Debug("CreateProjSectDescHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.CreateProjSectDescService(c.UserContext(), projSectDesc); err != nil {
			log.Error("CreateProjSectDescService error: ", err.Error())

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(model.SetResponse(fiber.StatusCreated, "created"))
	}
}

// @Summary Get projects section description by id.
// Description Get projects section description by id.
// @Tags projects section description
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /projects-section-description/{id} [get].
func GetProjSectDescByIDHandler(s ProjSectDescServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := struct {
			ID string `params:"id" validate:"required,uuid"`
		}{}

		if err := c.ParamsParser(&param); err != nil {
			log.Debug("GetProjSectDescByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(param); err != nil {
			log.Debug("GetProjSectDescByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		result, err := s.GetProjSectDescByIDService(c.UserContext(), param.ID)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				log.Debug("GetProjSectDescService error: ", err.Error())

				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}

			log.Error("GetProjSectDescService error: ", err.Error())

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// @Summary Update projects section description by id.
// Description Updates projects section description by id.
// @Tags projects section description
// @Accept json
// @Produce json
// @Param ProjSectDesc body model.UpdateProjSectDescSwagger true "ProjSectDesc"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /projects-section-description [put].
func UpdateProjSectDescByIDHandler(s ProjSectDescServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		projSectDesc := model.ProjSectDesc{}

		if err := c.BodyParser(&projSectDesc); err != nil {
			log.Debug("UpdateProjSectDescByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(projSectDesc); err != nil {
			log.Debug("UpdateProjSectDescByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.UpdateProjSectDescByIDService(c.UserContext(), projSectDesc); err != nil {
			if errors.Is(err, model.ErrNotFound) {
				log.Debug("UpdateProjSectDescByIDService error: ", err.Error())

				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}

			log.Error("UpdateProjSectDescByIDService error: ", err.Error())

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}
