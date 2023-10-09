package model

type CreatePartnerSwagger struct {
	Title string `json:"title" example:""`
	Image string `json:"image" example:""`
}

type UpdatePartnerSwagger struct {
	ID      string `json:"id" example:""`
	Created string `json:"created" example:""`
	CreatePartnerSwagger
}

type Partner struct {
	ID      string `json:"id" bson:"_id" validate:"omitempty,uuid"`
	Title   string `json:"title" bson:"title" validate:"required"`
	Image   string `json:"image" bson:"image" validate:"required"`
	Created string `json:"created" bson:"created"`
}

type PartnerQuery struct {
	Limit  int `query:"limit" validate:"lte=100"`
	Offset int `query:"offset"`
}
