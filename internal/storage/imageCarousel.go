package storage

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s Storage) CreateImagesCarouselStorage(ctx context.Context, imagesCarousel model.ImageCarousel) error {
	collection := s.DB.Collection(imagesCarouselCollection)

	_, err := collection.InsertOne(ctx, imagesCarousel)
	if err != nil {
		return handleError("error occurred in InsertOne", err)
	}

	return nil
}

func (s Storage) GetAllImagesCarouselStorage(ctx context.Context, query model.ImageCarouselQuery) ([]model.ImageCarousel, error) {
	collection := s.DB.Collection(imagesCarouselCollection)

	result := make([]model.ImageCarousel, 0)

	cursor, err := collection.Find(ctx, bson.D{}, limitAndOffset(query.Limit, query.Offset))
	if err != nil {
		return nil, handleError("error occurred in Find", err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		record := model.ImageCarousel{}

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

func (s Storage) GetImagesCarouselByIDStorage(ctx context.Context, id string) (model.ImageCarousel, error) {
	collection := s.DB.Collection(imagesCarouselCollection)

	imagesCarousel := model.ImageCarousel{}

	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&imagesCarousel); err != nil {
		return model.ImageCarousel{}, handleError("error occurred in FindOne", err)
	}

	return imagesCarousel, nil
}

func (s Storage) DeleteImagesCarouselByIDStorage(ctx context.Context, id string) error {
	collection := s.DB.Collection(imagesCarouselCollection)

	result, err := collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})

	return handleDeleteByIDError(result, "error occurred in DeleteOne", err)
}
