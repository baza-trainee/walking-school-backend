package storage

import (
	"context"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s Storage) CreateProjSectDescStorage(ctx context.Context, projSectDesc model.ProjSectDesc) error {
	collection := s.DB.Collection(projSectDescCollection)

	_, err := collection.InsertOne(ctx, projSectDesc)
	if err != nil {
		return handleError("error occurred in InsertOne", err)
	}

	return nil
}

func (s Storage) GetProjSectDescByIDStorage(ctx context.Context, id string) (model.ProjSectDesc, error) {
	collection := s.DB.Collection(projSectDescCollection)

	projSectDesc := model.ProjSectDesc{}

	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&projSectDesc); err != nil {
		return model.ProjSectDesc{}, handleError("error occurred in FindOne", err)
	}

	return projSectDesc, nil
}

func (s Storage) UpdateProjSectDescByIDStorage(ctx context.Context, projSectDesc model.ProjSectDesc) error {
	collection := s.DB.Collection(projSectDescCollection)

	result, err := collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: projSectDesc.ID}}, projSectDesc)

	return handleUpdateByIDError(result, "error occurred in ReplaceOne", err)
}
