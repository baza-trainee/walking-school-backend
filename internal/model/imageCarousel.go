package model

type CreateImageCarouselSwagger struct {
	Image string `json:"image" example:""`
}

type UpdateImageCarouselSwagger struct {
	ID string `json:"id" example:""`
	CreateImageCarouselSwagger
}

type ImageCarousel struct {
	ID    string `json:"id" bson:"_id" validate:"omitempty,uuid"`
	Image string `json:"image" bson:"image" validate:"required"`
}

type ImageCarouselQuery struct {
	Limit  int `query:"limit" validate:"lte=100"`
	Offset int `query:"offset"`
}
