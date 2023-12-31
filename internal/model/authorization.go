package model

import "github.com/golang-jwt/jwt"

const (
	AccessCookieName  = "access-cookie"
	RefreshCookieName = "refresh-cookie"
	AccessCookiePath  = "/api/v1"
	RefreshCookiePath = "/api/v1/authorization-refresh"
)

type Identity struct {
	Login    string `json:"login" bson:"login" validate:"email" example:"admin@example.com"`
	Password string `json:"password" bson:"password" validate:"min=10,max=255" example:"password777"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token" example:""`
	RefreshToken string `json:"refresh_token" example:""`
}

type Admin struct {
	ID       string `json:"id" bson:"_id" validate:"omitempty,uuid" example:""`
	Login    string `json:"login" bson:"login" validate:"email" example:"admin@example.com"`
	Password string `json:"password" bson:"password" validate:"min=10,max=255" example:"password777"`
}

type Claims struct {
	jwt.StandardClaims
	ID string
}

type Login struct {
	Login string `json:"login" bson:"login" validate:"email" example:""`
}

type ResetPassword struct {
	Token                string `json:"token" validate:"required" example:""`
	NewPassword          string `json:"new_password" validate:"min=10,max=255" example:"password888"`
	ConfirmedNewPassword string `json:"confirmed_new_password" validate:"min=10,max=255" example:"password888"`
}
