package storage

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s Storage) CreateUsertStorage(ctx context.Context, user model.User) error {
	collection := s.DB.Collection(userCollection)

	if err := collection.FindOne(ctx, creationFilter(user.Phone, user.Email)).Decode(&model.User{}); err == nil {
		return fmt.Errorf("may be such contacts already exist: %w", model.ErrConflict)
	}

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return handleError("error occurred in InsertOne", err)
	}

	return nil
}

func (s Storage) GetAllUserStorage(ctx context.Context, query model.UserQuery) ([]model.User, error) {
	collection := s.DB.Collection(userCollection)

	result := make([]model.User, 0)

	cursor, err := collection.Find(ctx, bson.D{}, limitAndOffset(query.Limit, query.Offset))
	if err != nil {
		return nil, handleError("error occurred in Find", err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		record := model.User{}

		if err := cursor.Decode(&record); err != nil {
			return nil, fmt.Errorf("error occurred in cursor.Decode: %w", err)
		}

		result = append(result, record)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error occurred in cursor.Err: %w", err)
	}

	return result, nil

}

func (s Storage) GetUserByIDStorage(ctx context.Context, id string) (model.User, error) {
	collection := s.DB.Collection(userCollection)

	user := model.User{}

	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&user); err != nil {
		return model.User{}, handleError("error occurred in FindOne", err)
	}

	return user, nil
}

func (s Storage) UpdateUserByIDStorage(ctx context.Context, user model.User) error {
	collection := s.DB.Collection(userCollection)

	if err := collection.FindOne(ctx, updateFilter(user.ID, user.Phone, user.Email)).Decode(&model.User{}); err == nil {
		return fmt.Errorf("may be such contacts already exist: %w", model.ErrConflict)
	}

	result, err := collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: user.ID}}, user)

	return handleUpdateByIDError(result, "error occurred in ReplaceOne", err)
}

func (s Storage) DeleteUserByIDStorage(ctx context.Context, id string) error {
	collection := s.DB.Collection(userCollection)

	result, err := collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})

	return handleDeleteByIDError(result, "error occurred in DeleteOne", err)
}
