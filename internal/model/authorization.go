package model

type Identity struct {
	Login    string `json:"login" bson:"login" validate:"email" example:"admin@example.com"`
	Password string `json:"password" bson:"password" validate:"min=6,max=255" example:"password123456"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Admin struct {
	ID string `json:"id" bson:"_id" validate:"uuid" example:""`
	Identity
}
