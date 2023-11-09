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
	FindAdminByID(context.Context, string) error
	FindAdminByLogin(context.Context, string) (model.Admin, error)
	ResetPasswordByID(context.Context, string, string) error
}

type Authorization struct {
	Storage AuthorizationStorageInterface
	CfgAuth config.AuthConfig
	CfgMsg  config.Feedback
}

func (a Authorization) SignInService(ctx context.Context, person model.Identity) (model.TokenPair, error) {
	passwordHash := SHA256(person.Password, a.CfgAuth.Salt)

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

func (a Authorization) RefreshService(ctx context.Context, token string) (model.TokenPair, error) {
	claim, err := ParseToken(token, a.CfgAuth.SigningKey)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("error occurred in ParseToken: %w", err)
	}

	if err := a.Storage.FindAdminByID(ctx, claim.ID); err != nil {
		return model.TokenPair{}, fmt.Errorf("error occurred in FindAdminByID: %w", err)
	}

	tokenPair, err := a.generateTokenPair(ctx, claim.ID)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("generateTokenPair error: %w", err)
	}

	return tokenPair, nil
}

func (a Authorization) ForgotPasswordService(ctx context.Context, login string) error {
	admin, err := a.Storage.FindAdminByLogin(ctx, login)
	if err != nil {
		return fmt.Errorf("error occurred in FindAdminByLogin: %w", err)
	}

	accessExpire := time.Now().Add(a.CfgAuth.AccessTokenTTL)

	token, err := generateJWT(accessExpire, a.CfgAuth.SigningKey, admin.ID)
	if err != nil {
		return fmt.Errorf("error occurred in generateJWT: %w", err)
	}

	link := fmt.Sprintf(linkToResetPassword, token)

	if err := sendMessage(
		a.CfgMsg.Host,
		a.CfgMsg.Port,
		a.CfgMsg.Username,
		a.CfgMsg.Password,
		a.CfgMsg.From,
		login,
		fmt.Sprintf(resetPasswordMessage, link),
	); err != nil {
		return fmt.Errorf("error occurred in sendMessage(): %w", err)
	}

	return nil
}

func (a Authorization) ResetPasswordService(ctx context.Context, data model.ResetPassword) error {
	claims, err := ParseToken(data.Token, a.CfgAuth.SigningKey)
	if err != nil {
		return fmt.Errorf("error occurred in ParseToken: %w", err)
	}

	passwordHash := SHA256(data.NewPassword, a.CfgAuth.Salt)

	if err := a.Storage.ResetPasswordByID(ctx, claims.ID, passwordHash); err != nil {
		return fmt.Errorf("error occurred in ResetPasswordByID: %w", err)
	}

	return nil
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

func (a Authorization) generateTokenPair(ctx context.Context, id string) (model.TokenPair, error) {
	accessExpire := time.Now().Add(a.CfgAuth.AccessTokenTTL)
	refreshExpire := time.Now().Add(a.CfgAuth.RefreshTokenTTL)

	accessToken, err := generateJWT(accessExpire, a.CfgAuth.SigningKey, id)
	if err != nil {
		return model.TokenPair{}, fmt.Errorf("generateJWT error: %w", err)
	}

	refreshToken, err := generateJWT(refreshExpire, a.CfgAuth.SigningKey, id)
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
