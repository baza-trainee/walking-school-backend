package model

type CreateProjSectDescSwagger struct {
	Description string `json:"description" example:"some description"`
}

type UpdateProjSectDescSwagger struct {
	ID string `json:"id" example:""`
	CreateProjSectDescSwagger
}

type ProjSectDesc struct {
	ID          string `json:"id" bson:"_id" validate:"omitempty,uuid"`
	Description string `json:"description" bson:"description" validate:"required"`
}
