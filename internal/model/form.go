package model

type CreateFormSwagger struct {
	Name    string `json:"" example:""`
	Surname string `json:"" example:""`
	Email   string `json:"" example:""`
	Phone   string `json:"" example:""`
	Text    string `json:"" example:""`
}

type Form struct {
	ID      string `json:"id" bson:"_id"`
	Name    string `json:"name" bson:"name" validate:"required"`
	Surname string `json:"surname" bson:"surname" validate:"required"`
	Email   string `json:"email" bson:"email" validate:"required,email"`
	Phone   string `json:"phone" bson:"phone" validate:"omitempty,e164,min=13,max=13"`
	Text    string `json:"text" bson:"text"`
}
