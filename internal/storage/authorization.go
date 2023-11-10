package storage

import (
	"context"
	"fmt"

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

func (s Storage) FindAdminByLogin(ctx context.Context, login string) (model.Admin, error) {
	collection := s.DB.Collection(adminCollection)

	admin := model.Admin{}

	if err := collection.FindOne(ctx, bson.D{{Key: "login", Value: login}}).Decode(&admin); err != nil {
		return model.Admin{}, handleError("error occurred in FindOne", err)
	}

	return admin, nil
}

func (s Storage) ResetPasswordByID(ctx context.Context, id, newPassword string) error {
	collection := s.DB.Collection(adminCollection)

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: newPassword}}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	return handleUpdateByIDError(result, "error occurred in UpdateOne()", err)
}

func (s Storage) RegistrationForTestStorage(ctx context.Context, admin model.Admin) error {
	collection := s.DB.Collection(adminCollection)

	_, err := s.FindAdminByLogin(ctx, admin.Login)
	if err == nil {
		return fmt.Errorf("such login is already exists: %w", model.ErrConflict)
	}

	_, err = collection.InsertOne(ctx, admin)
	if err != nil {
		return handleError("error occurred in InsertOne", err)
	}

	return nil
}
