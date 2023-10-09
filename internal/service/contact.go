package service

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/google/uuid"
)

type ContactStorageInterface interface {
	CreateContactStorage(context.Context, model.Contact) error
	GetContactByIDStorage(context.Context, string) (model.Contact, error)
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

func (c Contact) GetContactByIDService(ctx context.Context, param string) (model.Contact, error) {
	result, err := c.Storage.GetContactByIDStorage(ctx, param)
	if err != nil {
		return model.Contact{}, fmt.Errorf("error occurred in GetContactByIDStorage: %w", err)
	}

	return result, nil
}

func (c Contact) UpdateContactByIDService(ctx context.Context, contact model.Contact) error {
	if err := c.Storage.UpdateContactByIDStorage(ctx, contact); err != nil {
		return fmt.Errorf("error occurred in UpdateContactByIDStorage: %w", err)
	}

	return nil
}
