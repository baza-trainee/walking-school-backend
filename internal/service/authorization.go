package service

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/config"
	"github.com/baza-trainee/walking-school-backend/internal/model"
)

type AuthorizationStorageInterface interface {
	FindAdmin(context.Context, string, string) (model.Admin, error)
}

type Authorization struct {
	Storage AuthorizationStorageInterface
	Cfg     config.AuthConfig
}

func (a Authorization) SignInService(ctx context.Context, person model.Identity) (model.TokenPair, error) {
	passwordHash := SHA256(person.Password, a.Cfg.Salt)

	admin, err := a.Storage.FindAdmin(ctx, person.Login, passwordHash)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("error occurred in FindAdmin: %w", err)
	}

	admin.Password = ""

	tokenPair, err := a.generateTokenPair(ctx, admin.ID)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("generateTokenPair error: %w", err)
	}

	// return result, nil
}
