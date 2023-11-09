package model

import "github.com/golang-jwt/jwt"

const (
	AccessCookieName  = "access-cookie"
	RefreshCookieName = "refresh-cookie"
	AccessCookiePath  = "/"
	RefreshCookiePath = "/api/v1/authorization-refresh"
)

type Identity struct {
	Login    string `json:"login" bson:"login" validate:"email" example:"admin@example.com"`
	Password string `json:"password" bson:"password" validate:"min=6,max=255" example:"password777"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token" example:""`
	RefreshToken string `json:"refresh_token" example:""`
}

type Admin struct {
	ID string `json:"id" bson:"_id" validate:"uuid" example:""`
	Identity
}

type Claims struct {
	jwt.StandardClaims
	ID string
}

type Login struct {
	Login string `json:"login" bson:"login" validate:"email" example:""`
}
