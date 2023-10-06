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

type User struct {
	ID                        string   `json:"id,omitempty" bson:"_id" validate:"omitempty,uuid" example:""`
	Name                      string   `json:"name" bson:"name" validate:"required" example:""`
	Surname                   string   `json:"surname" bson:"surname" validate:"required" example:""`
	Patronymic                string   `json:"patronymic" bson:"patronymic" validate:"required" example:""`
	Location                  string   `json:"location" bson:"location" validate:"required" example:""`
	Phone                     string   `json:"phone" bson:"phone" validate:"omitempty,e164,min=13,max=13" example:""`
	Email                     string   `json:"email" bson:"email" validate:"omitempty,email" example:""`
	CombatCertificate         bool     `json:"combat_certificate" bson:"combat_certificate" validate:"required" example:"true"`
	DisabilityCertificate     []string `json:"disability_certificate" bson:"disability_certificate" validate:"required" example:""`
	InternationalPassport     []string `json:"international_passport" bson:"international_passport" validate:"required" example:""`
	WeightBelow95             bool     `json:"weight_below_95" bson:"weight_below_95" validate:"required" example:"true"`
	PhysicalActionConstraints []string `json:"physical_action_constraints" bson:"physical_action_constraints" validate:"required" example:""`
}

type UserQuery struct {
	Limit  int `query:"limit" validate:"lte=100"`
	Offset int `query:"offset"`
}
