package handler

import (
	"context"
	"log/slog"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type ProjSectDescServiceInterface interface {
	CreateProjSectDescService(context.Context, model.ProjSectDesc) error
	GetAllProjSectDescService(context.Context) ([]model.ProjSectDesc, error)
	UpdateProjSectDescByIDService(context.Context, model.ProjSectDesc) error
}

// @Summary Create project section description.
// Description Create project section description.
// @Tags project section description
// @Accept json
// @Produce json
// @Param ProjSectDesc body model.CreateProjSectDescSwagger true "ProjSectDesc"
// @Success 201 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /project-section-description [post].
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
			return handleError(log, "CreateProjSectDescService error: ", err)
		}

		return c.Status(fiber.StatusCreated).JSON(model.SetResponse(fiber.StatusCreated, "created"))
	}
}

// @Summary Get project section description.
// Description Get project section description.
// @Tags project section description
// @Produce json
// @Success 200 {object} model.Response
// @Success 204 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /project-section-description [get].
func GetAllProjSectDescHandler(s ProjSectDescServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := s.GetAllProjSectDescService(c.UserContext())
		if err != nil {
			return handleError(log, "GetAllProjSectDescService error: ", err)
		}

		if len(result) < minimalResult {
			return c.Status(fiber.StatusNoContent).JSON(model.SetResponse(fiber.StatusNoContent, "no content"))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// @Summary Update project section description by id.
// Description Updates project section description by id.
// @Tags project section description
// @Accept json
// @Produce json
// @Param ProjSectDesc body model.UpdateProjSectDescSwagger true "ProjSectDesc"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /project-section-description [put].
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
			return handleError(log, "UpdateProjSectDescByIDService error: ", err)
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}
