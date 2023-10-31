package service

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/google/uuid"
)

type ProjSectDescStorageInterface interface {
	CreateProjSectDescStorage(context.Context, model.ProjSectDesc) error
	GetAllProjSectDescStorage(context.Context) ([]model.ProjSectDesc, error)
	UpdateProjSectDescByIDStorage(context.Context, model.ProjSectDesc) error
}

type ProjSectDesc struct {
	Storage ProjSectDescStorageInterface
}

func (psd ProjSectDesc) CreateProjSectDescService(ctx context.Context, projSectDesc model.ProjSectDesc) error {
	projSectDesc.ID = uuid.NewString()

	if err := psd.Storage.CreateProjSectDescStorage(ctx, projSectDesc); err != nil {
		return fmt.Errorf("error occurred in CreateProjSectDescStorage: %w", err)
	}

	return nil
}

func (psd ProjSectDesc) GetAllProjSectDescService(ctx context.Context) ([]model.ProjSectDesc, error) {
	result, err := psd.Storage.GetAllProjSectDescStorage(ctx)
	if err != nil {
		return nil, fmt.Errorf("error occurred in GetAllProjSectDescStorage: %w", err)
	}

	return result, nil
}

func (psd ProjSectDesc) UpdateProjSectDescByIDService(ctx context.Context, projSectDesc model.ProjSectDesc) error {
	if err := psd.Storage.UpdateProjSectDescByIDStorage(ctx, projSectDesc); err != nil {
		return fmt.Errorf("error occurred in UpdateProjSectDescByIDStorage: %w", err)
	}

	return nil
}
