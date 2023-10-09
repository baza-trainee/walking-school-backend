package storage

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s Storage) CreatePartnerStorage(ctx context.Context, partner model.Partner) error {
	collection := s.DB.Collection(partnerCollection)

	_, err := collection.InsertOne(ctx, partner)
	if err != nil {
		return handleError("error occurred in InsertOne", err)
	}

	return nil
}

func (s Storage) GetAllPartnerStorage(ctx context.Context, query model.PartnerQuery) ([]model.Partner, error) {
	collection := s.DB.Collection(partnerCollection)

	result := make([]model.Partner, 0)

	cursor, err := collection.Find(ctx, bson.D{}, limitAndOffset(query.Limit, query.Offset))
	if err != nil {
		return nil, handleError("error occurred in Find", err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		record := model.Partner{}

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

func (s Storage) GetPartnerByIDStorage(ctx context.Context, id string) (model.Partner, error) {
	collection := s.DB.Collection(partnerCollection)

	partner := model.Partner{}

	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&partner); err != nil {
		return model.Partner{}, handleError("error occurred in FindOne", err)
	}

	return partner, nil
}

func (s Storage) UpdatePartnerByIDStorage(ctx context.Context, partner model.Partner) error {
	collection := s.DB.Collection(partnerCollection)

	result, err := collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: partner.ID}}, partner)

	return handleUpdateByIDError(result, "error occurred in ReplaceOne", err)
}

func (s Storage) DeletePartnerByIDStorage(ctx context.Context, id string) error {
	collection := s.DB.Collection(partnerCollection)

	result, err := collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})

	return handleDeleteByIDError(result, "error occurred in DeleteOne", err)
}
