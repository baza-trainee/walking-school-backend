package service

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/google/uuid"
)

type HeroStorageInterface interface {
	CreateHeroStorage(context.Context, model.Hero) error
	GetAllHeroStorage(context.Context, model.HeroQuery) ([]model.Hero, error)
	GetHeroByIDStorage(context.Context, string) (model.Hero, error)
	UpdateHeroByIDStorage(context.Context, model.Hero) error
	DeleteHeroByIDStorage(context.Context, string) error
}

type Hero struct {
	Storage HeroStorageInterface
}

func (h Hero) CreateHeroService(ctx context.Context, hero model.Hero) error {
	hero.ID = uuid.NewString()

	if err := h.Storage.CreateHeroStorage(ctx, hero); err != nil {
		return fmt.Errorf("error occurred in CreateHeroStorage: %w", err)
	}

	return nil
}

func (h Hero) GetAllHeroService(ctx context.Context, query model.HeroQuery) ([]model.Hero, error) {
	result, err := h.Storage.GetAllHeroStorage(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error occurred in GetAllHeroStorage: %w", err)
	}

	return result, nil
}

func (h Hero) GetHeroByIDService(ctx context.Context, param string) (model.Hero, error) {
	result, err := h.Storage.GetHeroByIDStorage(ctx, param)
	if err != nil {
		return model.Hero{}, fmt.Errorf("error occurred in GetHeroByIDStorage: %w", err)
	}

	return result, nil
}

func (h Hero) UpdateHeroByIDService(ctx context.Context, hero model.Hero) error {
	if err := h.Storage.UpdateHeroByIDStorage(ctx, hero); err != nil {
		return fmt.Errorf("error occurred in UpdateHeroByIDStorage: %w", err)
	}

	return nil
}

func (h Hero) DeleteHeroByIDService(ctx context.Context, param string) error {
	if err := h.Storage.DeleteHeroByIDStorage(ctx, param); err != nil {
		return fmt.Errorf("error occurred in DeleteHeroByIDStorage: %w", err)
	}

	return nil
}
