package service

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/google/uuid"
)

type UserStorageInterface interface {
	CreateUsertStorage(context.Context, model.User) error
	GetAllUserStorage(context.Context, model.UserQuery) ([]model.User, error)
	GetUserByIDStorage(context.Context, string) (model.User, error)
	UpdateUserByIDStorage(context.Context, model.User) error
	DeleteUserByIDStorage(context.Context, string) error
}

type User struct {
	Storage UserStorageInterface
}

func (u User) CreateUserService(ctx context.Context, user model.User) error {
	user.ID = uuid.NewString()

	if err := u.Storage.CreateUsertStorage(ctx, user); err != nil {
		return fmt.Errorf("error occurred in CreateUsertStorage: %w", err)
	}

	return nil
}

func (u User) GetAllUserService(ctx context.Context, query model.UserQuery) ([]model.User, error) {
	result, err := u.Storage.GetAllUserStorage(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error occurred in GetAllUserStorage: %w", err)
	}

	return result, nil
}

func (u User) GetUserByIDService(ctx context.Context, param string) (model.User, error) {
	result, err := u.Storage.GetUserByIDStorage(ctx, param)
	if err != nil {
		return model.User{}, fmt.Errorf("error occurred in GetUserByIDStorage: %w", err)
	}

	return result, nil
}

func (u User) UpdateUserByIDService(ctx context.Context, user model.User) error {
	if err := u.Storage.UpdateUserByIDStorage(ctx, user); err != nil {
		return fmt.Errorf("error occurred in UpdateUserByIDStorage: %w", err)
	}

	return nil
}

func (u User) DeleteUserByIDService(ctx context.Context, param string) error {
	if err := u.Storage.DeleteUserByIDStorage(ctx, param); err != nil {
		return fmt.Errorf("error occurred in DeleteUserByIDStorage: %w", err)
	}

	return nil
}
