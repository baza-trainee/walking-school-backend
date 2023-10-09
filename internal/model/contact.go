package model

type CreateContactSwagger struct {
	Phone        string `json:"phone" example:""`
	ContactEmail string `json:"contact_email" example:""`
	AnswerEmail  string `json:"answer_email" example:""`
	Facebook     string `json:"facebook" example:""`
	LinkedIn     string `json:"linkedin" example:""`
	Telegram     string `json:"telegram" example:""`
}

type UpdateContactSwagger struct {
	ID string `json:"id" example:""`
	CreateContactSwagger
}

type Contact struct {
	ID           string `json:"id" bson:"_id" validate:"omitempty,uuid"`
	Phone        string `json:"phone" bson:"phone" validate:"required,e164,min=13,max=13"`
	ContactEmail string `json:"contact_email" bson:"contact_email" validate:"required,email"`
	AnswerEmail  string `json:"answer_email" bson:"answer_email" validate:"required,email"`
	Facebook     string `json:"facebook" bson:"facebook"`
	LinkedIn     string `json:"linkedin" bson:"linkedin"`
	Telegram     string `json:"telegram" bson:"telegram"`
}
