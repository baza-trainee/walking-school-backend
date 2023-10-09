package service

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/google/uuid"
)

type ProjSectDescStorageInterface interface {
	CreateProjSectDescStorage(context.Context, model.ProjSectDesc) error
	GetProjSectDescByIDStorage(context.Context, string) (model.ProjSectDesc, error)
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

func (psd ProjSectDesc) GetProjSectDescByIDService(ctx context.Context, param string) (model.ProjSectDesc, error) {
	result, err := psd.Storage.GetProjSectDescByIDStorage(ctx, param)
	if err != nil {
		return model.ProjSectDesc{}, fmt.Errorf("error occurred in GetProjSectDescByIDStorage: %w", err)
	}

	return result, nil
}

func (psd ProjSectDesc) UpdateProjSectDescByIDService(ctx context.Context, projSectDesc model.ProjSectDesc) error {
	if err := psd.Storage.UpdateProjSectDescByIDStorage(ctx, projSectDesc); err != nil {
		return fmt.Errorf("error occurred in UpdateProjSectDescByIDStorage: %w", err)
	}

	return nil
}
