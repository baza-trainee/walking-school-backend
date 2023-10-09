package storage

import (
	"context"

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

func (s Storage) GetContactByIDStorage(ctx context.Context, id string) (model.Contact, error) {
	collection := s.DB.Collection(contactCollection)

	contact := model.Contact{}

	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&contact); err != nil {
		return model.Contact{}, handleError("error occurred in FindOne", err)
	}

	return contact, nil
}

func (s Storage) UpdateContactByIDStorage(ctx context.Context, contact model.Contact) error {
	collection := s.DB.Collection(contactCollection)

	result, err := collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: contact.ID}}, contact)

	return handleUpdateByIDError(result, "error occurred in ReplaceOne", err)
}
