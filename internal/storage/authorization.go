package storage

import (
	"context"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s Storage) FindAdmin(ctx context.Context, login, password string) (model.Admin, error) {
	collection := s.DB.Collection(adminCollection)

	admin := model.Admin{}

	if err := collection.FindOne(ctx, bson.D{{Key: "$and", Value: bson.A{
		bson.D{{Key: "login", Value: login}},
		bson.D{{Key: "password", Value: password}},
	}}}).Decode(&admin); err != nil {
		return model.Admin{}, handleError("error occurred in FindOne", err)
	}

	return admin, nil
}

func (s Storage) FindAdminByID(ctx context.Context, id string) error {
	collection := s.DB.Collection(adminCollection)

	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&model.Admin{}); err != nil {
		return handleError("error occurred in FindOne", err)
	}

	return nil
}
