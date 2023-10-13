package model

type CreateImageCarouselSwagger struct {
	Image []string `json:"image" example:""`
}

type UpdateImageCarouselSwagger struct {
	ID string `json:"id" example:""`
	CreateImageCarouselSwagger
}

type ImageCarousel struct {
	ID    string   `json:"id,omitempty" bson:"_id" validate:"omitempty,uuid"`
	Image []string `json:"image" bson:"image" validate:"required,lte=6"`
}
