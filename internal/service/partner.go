package service

import (
	"context"
	"fmt"
	"time"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/google/uuid"
)

type PartnerStorageInterface interface {
	CreatePartnerStorage(context.Context, model.Partner) error
	GetAllPartnerStorage(context.Context, model.PartnerQuery) ([]model.Partner, error)
	GetPartnerByIDStorage(context.Context, string) (model.Partner, error)
	UpdatePartnerByIDStorage(context.Context, model.Partner) error
	DeletePartnerByIDStorage(context.Context, string) error
}

type Partner struct {
	Storage PartnerStorageInterface
}

func (p Partner) CreatePartnerService(ctx context.Context, partner model.Partner) error {
	partner.ID = uuid.NewString()
	partner.Created = time.Now().Format("01-2006")

	if err := p.Storage.CreatePartnerStorage(ctx, partner); err != nil {
		return fmt.Errorf("error occurred in CreatePartnerStorage: %w", err)
	}

	return nil
}

func (p Partner) GetAllPartnerService(ctx context.Context, query model.PartnerQuery) ([]model.Partner, error) {
	result, err := p.Storage.GetAllPartnerStorage(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error occurred in GetAllPartnerStorage: %w", err)
	}

	if len(result) < minimalResult {
		return []model.Partner{}, model.ErrNoContent
	}

	return result, nil
}

func (p Partner) GetPartnerByIDService(ctx context.Context, param string) (model.Partner, error) {
	result, err := p.Storage.GetPartnerByIDStorage(ctx, param)
	if err != nil {
		return model.Partner{}, fmt.Errorf("error occurred in GetPartnerByIDStorage: %w", err)
	}

	return result, nil
}

func (p Partner) UpdatePartnerByIDService(ctx context.Context, partner model.Partner) error {
	if err := p.Storage.UpdatePartnerByIDStorage(ctx, partner); err != nil {
		return fmt.Errorf("error occurred in UpdatePartnerByIDStorage: %w", err)
	}

	return nil
}

func (p Partner) DeletePartnerByIDService(ctx context.Context, param string) error {
	if err := p.Storage.DeletePartnerByIDStorage(ctx, param); err != nil {
		return fmt.Errorf("error occurred in DeletePartnerByIDStorage: %w", err)
	}

	return nil
}
