package storage

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s Storage) CreateContactStorage(ctx context.Context, contact model.Contact) error {
	collection := s.DB.Collection(contactCollection)

	_, err := collection.InsertOne(ctx, contact)
	if err != nil {
		return handleError("error occurred in InsertOne", err)
	}

	return nil
}

func (s Storage) GetAllContactStorage(ctx context.Context) ([]model.Contact, error) {
	collection := s.DB.Collection(contactCollection)

	result := make([]model.Contact, 0)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, handleError("error occurred in Find", err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		record := model.Contact{}

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

func (s Storage) UpdateContactByIDStorage(ctx context.Context, contact model.Contact) error {
	collection := s.DB.Collection(contactCollection)

	result, err := collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: contact.ID}}, contact)

	return handleUpdateByIDError(result, "error occurred in ReplaceOne", err)
}
