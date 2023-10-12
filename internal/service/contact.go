package service

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/google/uuid"
)

type ContactStorageInterface interface {
	CreateContactStorage(context.Context, model.Contact) error
	GetAllContactStorage(context.Context) ([]model.Contact, error)
	UpdateContactByIDStorage(context.Context, model.Contact) error
}

type Contact struct {
	Storage ContactStorageInterface
}

func (c Contact) CreateContactService(ctx context.Context, contact model.Contact) error {
	contact.ID = uuid.NewString()

	if err := c.Storage.CreateContactStorage(ctx, contact); err != nil {
		return fmt.Errorf("error occurred in CreateContactStorage: %w", err)
	}

	return nil
}

func (c Contact) GetAllContactService(ctx context.Context) ([]model.Contact, error) {
	result, err := c.Storage.GetAllContactStorage(ctx)
	if err != nil {
		return nil, fmt.Errorf("error occurred in GetAllContactStorage: %w", err)
	}

	if len(result) < minimalResult {
		return []model.Contact{}, model.ErrNoContent
	}

	return result, nil
}

func (c Contact) UpdateContactByIDService(ctx context.Context, contact model.Contact) error {
	if err := c.Storage.UpdateContactByIDStorage(ctx, contact); err != nil {
		return fmt.Errorf("error occurred in UpdateContactByIDStorage: %w", err)
	}

	return nil
}
