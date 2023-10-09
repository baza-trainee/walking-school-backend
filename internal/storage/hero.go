package storage

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s Storage) CreateHeroStorage(ctx context.Context, hero model.Hero) error {
	collection := s.DB.Collection(heroCollection)

	_, err := collection.InsertOne(ctx, hero)
	if err != nil {
		return handleError("error occurred in InsertOne", err)
	}

	return nil
}

func (s Storage) GetAllHeroStorage(ctx context.Context, query model.HeroQuery) ([]model.Hero, error) {
	collection := s.DB.Collection(heroCollection)

	result := make([]model.Hero, 0)

	cursor, err := collection.Find(ctx, bson.D{}, limitAndOffset(query.Limit, query.Offset))
	if err != nil {
		return nil, handleError("error occurred in Find", err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		record := model.Hero{}

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

func (s Storage) GetHeroByIDStorage(ctx context.Context, id string) (model.Hero, error) {
	collection := s.DB.Collection(heroCollection)

	hero := model.Hero{}

	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&hero); err != nil {
		return model.Hero{}, handleError("error occurred in FindOne", err)
	}

	return hero, nil
}

func (s Storage) UpdateHeroByIDStorage(ctx context.Context, hero model.Hero) error {
	collection := s.DB.Collection(heroCollection)

	result, err := collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: hero.ID}}, hero)

	return handleUpdateByIDError(result, "error occurred in ReplaceOne", err)
}

func (s Storage) DeleteHeroByIDStorage(ctx context.Context, id string) error {
	collection := s.DB.Collection(heroCollection)

	result, err := collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})

	return handleDeleteByIDError(result, "error occurred in DeleteOne", err)
}
