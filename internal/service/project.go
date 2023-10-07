package service

import (
	"context"
	"fmt"
	"time"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/google/uuid"
)

type ProjectStorageInterface interface {
	CreateProjectStorage(context.Context, model.Project) error
	GetAllProjectStorage(context.Context, model.ProjectQuery) ([]model.Project, error)
	GetProjectByIDStorage(context.Context, string) (model.Project, error)
	UpdateProjectByIDStorage(context.Context, model.Project) error
	DeleteProjectByIDStorage(context.Context, string) error
}

type Project struct {
	Storage ProjectStorageInterface
}

func (p Project) CreateProjectService(ctx context.Context, project model.Project) error {
	project.ID = uuid.NewString()
	project.Created = time.Now().Format("01-2006")
	project.LastModified = time.Now().Format("01-2006")
	project.IsActive = true

	if err := p.Storage.CreateProjectStorage(ctx, project); err != nil {
		return fmt.Errorf("error occurred in CreateProjectStorage: %w", err)
	}

	return nil
}

func (p Project) GetAllProjectService(ctx context.Context, query model.ProjectQuery) ([]model.Project, error) {
	result, err := p.Storage.GetAllProjectStorage(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error occurred in GetAllProjectStorage: %w", err)
	}

	if len(result) < minimalResult {
		return []model.Project{}, model.ErrNoContent
	}

	return result, nil
}

func (p Project) GetProjectByIDService(ctx context.Context, param string) (model.Project, error) {
	result, err := p.Storage.GetProjectByIDStorage(ctx, param)
	if err != nil {
		return model.Project{}, fmt.Errorf("error occurred in CreateProjectByIDStorage: %w", err)
	}

	return result, nil
}

func (p Project) UpdateProjectByIDService(ctx context.Context, project model.Project) error {
	if project.IsActive == true {
		project.LastModified = time.Now().Format("01-2006")
	}

	if err := p.Storage.UpdateProjectByIDStorage(ctx, project); err != nil {
		return fmt.Errorf("error occurred in UpdateProjectByIDStorage: %w", err)
	}

	return nil
}

func (p Project) DeleteProjectByIDService(ctx context.Context, param string) error {
	if err := p.Storage.DeleteProjectByIDStorage(ctx, param); err != nil {
		return fmt.Errorf("error occurred in DeleteProjectByIDStorage: %w", err)
	}

	return nil
}
