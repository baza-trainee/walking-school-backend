package model

type CreateHeroSwagger struct {
	Title       string `json:"title" example:"some title"`
	Description string `json:"description" example:"some description"`
	Image       string `json:"image" example:"stunning image"`
}

type UpdateHeroSwagger struct {
	ID string `json:"id" example:""`
	CreateHeroSwagger
}

type Hero struct {
	ID          string `json:"id" bson:"_id" validate:"omitempty,uuid"`
	Title       string `json:"title" bson:"title" validate:"required"`
	Description string `json:"description" bson:"description" validate:"required"`
	Image       string `json:"image" bson:"image" validate:"required"`
}

type HeroQuery struct {
	Limit  int `query:"limit" validate:"lte=100"`
	Offset int `query:"offset"`
}
