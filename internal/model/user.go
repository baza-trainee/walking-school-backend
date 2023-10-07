package model

type CreateUserSwagger struct {
	Name                      string   `json:"name" example:""`
	Surname                   string   `json:"surname" example:""`
	Patronymic                string   `json:"patronymic" example:""`
	Location                  string   `json:"location" example:""`
	Phone                     string   `json:"phone" example:"+380631122331"`
	Email                     string   `json:"email" example:"example1@gmail.com"`
	CombatCertificate         bool     `json:"combat_certificate" example:"true"`
	DisabilityCertificate     []string `json:"disability_certificate" example:""`
	InternationalPassport     []string `json:"international_passport" example:""`
	WeightBelow95             bool     `json:"weight_below_95" example:"true"`
	PhysicalActionConstraints []string `json:"physical_action_constraints" example:""`
}

type UpdateUserSwagger struct {
	ID string `json:"id,omitempty" example:""`
	CreateUserSwagger
}

type User struct {
	ID                        string   `json:"id,omitempty" bson:"_id" validate:"omitempty,uuid"`
	Name                      string   `json:"name" bson:"name" validate:"required"`
	Surname                   string   `json:"surname" bson:"surname" validate:"required"`
	Patronymic                string   `json:"patronymic" bson:"patronymic" validate:"required"`
	Location                  string   `json:"location" bson:"location" validate:"required"`
	Phone                     string   `json:"phone" bson:"phone" validate:"omitempty,e164,min=13,max=13"`
	Email                     string   `json:"email" bson:"email" validate:"omitempty,email"`
	CombatCertificate         bool     `json:"combat_certificate" bson:"combat_certificate" validate:"required"`
	DisabilityCertificate     []string `json:"disability_certificate" bson:"disability_certificate" validate:"required"`
	InternationalPassport     []string `json:"international_passport" bson:"international_passport" validate:"required"`
	WeightBelow95             bool     `json:"weight_below_95" bson:"weight_below_95" validate:"required"`
	PhysicalActionConstraints []string `json:"physical_action_constraints" bson:"physical_action_constraints" validate:"required"`
}

type UserQuery struct {
	Limit  int `query:"limit" validate:"lte=100"`
	Offset int `query:"offset"`
}
