package service

import (
	"context"
	"fmt"
	"time"

	"github.com/baza-trainee/walking-school-backend/internal/config"
	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/golang-jwt/jwt"
)

type AuthorizationStorageInterface interface {
	FindAdmin(context.Context, string, string) (model.Admin, error)
}

type Authorization struct {
	Storage AuthorizationStorageInterface
	Cfg     config.AuthConfig
}

func ParseToken(tokenString, signingKey string) (model.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, model.ErrInvalidSigningMethod
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return model.Claims{}, fmt.Errorf("accessToken throws an error during parsing: %w", err)
	}

	claims, ok := token.Claims.(*model.Claims)
	if !ok {
		return model.Claims{}, model.ErrWrongTokenClaimType
	}

	return *claims, nil
}

func (a Authorization) SignInService(ctx context.Context, person model.Identity) (model.TokenPair, error) {
	passwordHash := SHA256(person.Password, a.Cfg.Salt)

	admin, err := a.Storage.FindAdmin(ctx, person.Login, passwordHash)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("error occurred in FindAdmin: %w", err)
	}

	tokenPair, err := a.generateTokenPair(ctx, admin.ID)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("generateTokenPair error: %w", err)
	}

	return tokenPair, nil
}

func (a Authorization) generateTokenPair(ctx context.Context, id string) (model.TokenPair, error) {
	accessExpire := time.Now().Add(a.Cfg.AccessTokenTTL)
	refreshExpire := time.Now().Add(a.Cfg.RefreshTokenTTL)

	accessToken, err := generateJWT(accessExpire, a.Cfg.SigningKey, id)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("generateJWT error: %w", err)
	}

	refreshToken, err := generateJWT(refreshExpire, a.Cfg.SigningKey, id)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("generateJWT error: %w", err)
	}

	tokenPair := model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokenPair, nil

}

func generateJWT(expireToken time.Time, signingKey, id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		model.Claims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expireToken.Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			ID: id,
		})

	tokenValue, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", fmt.Errorf("token.SignedString error: %w", err)
	}

	return tokenValue, nil

}
