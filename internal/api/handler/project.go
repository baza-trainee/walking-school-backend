package handler

import (
	"context"
	"errors"
	"log/slog"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type ProjectServiceInterface interface {
	CreateProjectService(context.Context, model.Project) error
	GetAllProjectService(context.Context, model.ProjectQuery) ([]model.Project, error)
	GetProjectByIDService(context.Context, string) (model.Project, error)
	UpdateProjectByIDService(context.Context, model.Project) error
	DeleteProjectByIDService(context.Context, string) error
}

// @Summary Create project.
// Description Creates project.
// @Tags project
// @Accept json
// @Produce json
// @Param Project body model.CreateProjectSwagger true "Project"
// @Success 201 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /project [post].
func CreateProjectHandler(s ProjectServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		project := model.Project{}

		if err := c.BodyParser(&project); err != nil {
			log.Debug("CreateProjectHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(project); err != nil {
			log.Debug("CreateProjectHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.CreateProjectService(c.UserContext(), project); err != nil {
			log.Error("CreateProjectService error: ", err.Error())
			// Какие ошибки могут возвращаться?
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(model.SetResponse(fiber.StatusCreated, "created"))
	}
}

// @Summary Get all projects.
// Description Get all projects.
// @Tags project
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} model.Response
// @Success 204 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /project [get].
func GetAllProjectHandler(s ProjectServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := model.ProjectQuery{
			Limit:  standartLimitValue,
			Offset: standartOffsetValue,
		}

		if err := c.QueryParser(&query); err != nil {
			log.Debug("GetAllProjectHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(query); err != nil {
			log.Debug("GetAllProjectHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		result, err := s.GetAllProjectService(c.UserContext(), query)
		if err != nil {
			if errors.Is(err, model.ErrNoContent) {
				log.Debug("GetAllProjectService error: ", err.Error())

				return fiber.NewError(fiber.StatusNoContent, err.Error())
			}

			log.Error("GetAllProjectService error: ", err.Error())
			// Какие ошибки могут возвращаться?
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// @Summary Get project by id.
// Description Gets project by id.
// @Tags project
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /project/{id} [get].
func GetProjectByIDHandler(s ProjectServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := struct {
			ID string `params:"id" validate:"required,uuid"`
		}{}

		if err := c.ParamsParser(&param); err != nil {
			log.Debug("GetProjectByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(param); err != nil {
			log.Debug("GetProjectByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		result, err := s.GetProjectByIDService(c.UserContext(), param.ID)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				log.Debug("GetProjectService error: ", err.Error())

				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}

			log.Error("GetProjectService error: ", err.Error())
			// Какие ошибки могут возвращаться?
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// @Summary Update project by id.
// Description Updates project by id.
// @Tags project
// @Accept json
// @Produce json
// @Param Project body model.Project true "Project"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /project [put].
func UpdateProjectByIDHandler(s ProjectServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		project := model.Project{}

		if err := c.BodyParser(&project); err != nil {
			log.Debug("UpdateProjectByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(project); err != nil {
			log.Debug("UpdateProjectByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.UpdateProjectByIDService(c.UserContext(), project); err != nil {
			if errors.Is(err, model.ErrNotFound) {
				log.Debug("UpdateProjectByIDService error: ", err.Error())

				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}

			log.Error("UpdateProjectByIDService error: ", err.Error())
			// Какие ошибки могут возвращаться?
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}

// @Summary Delete project by id.
// Description Deletes project by id.
// @Tags project
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 408 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /project/{id} [delete].
func DeleteProjectByIDHandler(s ProjectServiceInterface, log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := struct {
			ID string `params:"id" validate:"required,uuid"`
		}{}

		if err := c.ParamsParser(&param); err != nil {
			log.Debug("DeleteProjectByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(param); err != nil {
			log.Debug("DeleteProjectByIDHandler error: ", err.Error())

			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := s.DeleteProjectByIDService(c.UserContext(), param.ID); err != nil {
			if errors.Is(err, model.ErrNotFound) {
				log.Debug("DeleteProjectByIDService error: ", err.Error())

				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}

			log.Error("DeleteProjectByIDService error: ", err.Error())
			// Какие ошибки могут возвращаться?
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(model.SetResponse(fiber.StatusOK, "success"))
	}
}
