package storage

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s Storage) CreateUsertStorage(ctx context.Context, user model.User) error {
	collection := s.DB.Collection(userCollection)

	// decode := model.User{}

	if err := collection.FindOne(ctx, contactFilter(user.Phone, user.Email)).Decode(&model.User{}); err == nil {
		// fmt.Println("-------------------------------")
		// fmt.Println(user)
		// fmt.Println("------/-/-/-/-/-/-/-/-/-/-/-/-/-/------------")
		// fmt.Println(decode)
		// fmt.Println("-------------------------------")
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
