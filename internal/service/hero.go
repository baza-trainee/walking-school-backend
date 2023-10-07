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

func (p Hero) CreateHeroService(ctx context.Context, hero model.Hero) error {
	hero.ID = uuid.NewString()

	if err := p.Storage.CreateHeroStorage(ctx, hero); err != nil {
		return fmt.Errorf("error occurred in CreateHeroStorage: %w", err)
	}

	return nil
}

func (p Hero) GetAllHeroService(ctx context.Context, query model.HeroQuery) ([]model.Hero, error) {
	result, err := p.Storage.GetAllHeroStorage(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error occurred in GetAllHeroStorage: %w", err)
	}

	if len(result) < minimalResult {
		return []model.Hero{}, model.ErrNoContent
	}

	return result, nil
}

func (p Hero) GetHeroByIDService(ctx context.Context, param string) (model.Hero, error) {
	result, err := p.Storage.GetHeroByIDStorage(ctx, param)
	if err != nil {
		return model.Hero{}, fmt.Errorf("error occurred in GetHeroByIDStorage: %w", err)
	}

	return result, nil
}

func (p Hero) UpdateHeroByIDService(ctx context.Context, hero model.Hero) error {
	if err := p.Storage.UpdateHeroByIDStorage(ctx, hero); err != nil {
		return fmt.Errorf("error occurred in UpdateHeroByIDStorage: %w", err)
	}

	return nil
}

func (p Hero) DeleteHeroByIDService(ctx context.Context, param string) error {
	if err := p.Storage.DeleteHeroByIDStorage(ctx, param); err != nil {
		return fmt.Errorf("error occurred in DeleteHeroByIDStorage: %w", err)
	}

	return nil
}
