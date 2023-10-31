package storage

import (
	"context"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s Storage) FindAdmin(ctx context.Context, login, password string) (model.Admin, error) {
	collection := s.DB.Collection(adminCollection)

	admin := model.Admin{}

	if err := collection.FindOne(ctx, bson.E{Key: "$and", Value: bson.A{
		bson.E{Key: "login", Value: login},
		bson.E{Key: "password", Value: password},
	}}).Decode(&admin); err != nil {
		return model.Admin{}, handleError("error occurred in FindOne", err)
	}

	return admin, nil
}
