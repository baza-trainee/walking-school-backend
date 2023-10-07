package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s Storage) CreateProjSectDescStorage(ctx context.Context, projSectDesc model.ProjSectDesc) error {
	collection := s.DB.Collection(projSectDescCollection)

	_, err := collection.InsertOne(ctx, projSectDesc)
	if err != nil {

		return fmt.Errorf("error occurred in InsertOne: %w", err)
	}

	return nil
}

func (s Storage) GetProjSectDescByIDStorage(ctx context.Context, id string) (model.ProjSectDesc, error) {
	collection := s.DB.Collection(projSectDescCollection)

	projSectDesc := model.ProjSectDesc{}

	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&projSectDesc); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.ProjSectDesc{}, model.ErrNotFound
		}

		return model.ProjSectDesc{}, fmt.Errorf("error occurred in FindOne: %w", err)
	}

	return projSectDesc, nil
}

func (s Storage) UpdateProjSectDescByIDStorage(ctx context.Context, projSectDesc model.ProjSectDesc) error {
	collection := s.DB.Collection(projSectDescCollection)

	result, err := collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: projSectDesc.ID}}, projSectDesc)
	if err != nil {
		return fmt.Errorf("error occurred in ReplaceOne: %w", err)
	}

	if result.MatchedCount != matchedOneDocument {
		return model.ErrNotFound
	}

	return nil
}
