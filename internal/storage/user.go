package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s Storage) CreateUsertStorage(ctx context.Context, user model.User) error {
	collection := s.DB.Collection(userCollection)

	if err := collection.FindOne(ctx, creationFilter(user.Phone, user.Email)).Decode(&model.User{}); err == nil {
		return fmt.Errorf("may be such contacts already exist: %w", model.ErrConflict)
	}

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		// Какие ошибки могут возвращаться?
		return fmt.Errorf("error occurred in InsertOne: %w", err)
	}

	return nil
}

func (s Storage) GetAllUserStorage(ctx context.Context, query model.UserQuery) ([]model.User, error) {
	collection := s.DB.Collection(userCollection)

	result := make([]model.User, 0)

	cursor, err := collection.Find(ctx, bson.D{}, LimitAndOffset(query.Limit, query.Offset))
	if err != nil {
		// Какие ошибки могут возвращаться?
		return nil, fmt.Errorf("error occurred in Find: %w", err)
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
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.User{}, model.ErrNotFound
		}

		return model.User{}, fmt.Errorf("error occurred in FindOne: %w", err)
	}

	return user, nil
}

func (s Storage) UpdateUserByIDStorage(ctx context.Context, user model.User) error {
	collection := s.DB.Collection(userCollection)

	if err := collection.FindOne(ctx, updateFilter(user.ID, user.Phone, user.Email)).Decode(&model.User{}); err == nil {
		return fmt.Errorf("may be such contacts already exist: %w", model.ErrConflict)
	}

	result, err := collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: user.ID}}, user)
	if err != nil {
		return fmt.Errorf("error occurred in ReplaceOne: %w", err)
	}

	if result.MatchedCount != matchedOneDocument {
		return model.ErrNotFound
	}

	return nil
}
