package model

type CreateProjectSwagger struct {
	Title       string   `json:"title" example:""`
	Description string   `json:"description" example:""`
	Link        string   `json:"link" example:""`
	Image       string   `json:"image" example:""`
	Period      []string `json:"period" example:""`
	Category    string   `json:"category" example:""`
	AgeCategory string   `json:"age_category" example:""`
}

type Project struct {
	ID           string   `json:"id,omitempty" bson:"_id" validate:"omitempty,uuid" example:""`
	Title        string   `json:"title" bson:"title" validate:"required" example:""`
	Description  string   `json:"description" bson:"description" validate:"required" example:""`
	Link         string   `json:"link" bson:"link" validate:"required" example:""`
	Date         string   `json:"date" bson:"date" example:""`
	LastModified string   `json:"last_modified" bson:"last_modified" example:""`
	Image        string   `json:"image" bson:"image" validate:"required" example:""`
	IsActive     bool     `json:"is_active" bson:"is_active" example:"true"`
	Period       []string `json:"period" bson:"period" validate:"required,len=2" example:""`
	Category     string   `json:"category" bson:"category" validate:"required" example:""`
	AgeCategory  string   `json:"age_category" bson:"age_category" validate:"required" example:""`
}

type ProjectQuery struct {
	Limit  int `query:"limit" validate:"lte=100"`
	Offset int `query:"offset"`
}
