package storage

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s Storage) CreateProjectStorage(ctx context.Context, project model.Project) error {
	collection := s.DB.Collection(projectCollection)

	_, err := collection.InsertOne(ctx, project)
	if err != nil {
		return handleError("error occurred in InsertOne", err)
	}

	return nil
}

func (s Storage) GetAllProjectStorage(ctx context.Context, query model.ProjectQuery) ([]model.Project, error) {
	collection := s.DB.Collection(projectCollection)

	result := make([]model.Project, 0)

	cursor, err := collection.Find(ctx, bson.D{}, limitAndOffset(query.Limit, query.Offset))
	if err != nil {
		return nil, handleError("error occurred in Find", err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		record := model.Project{}

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

func (s Storage) GetProjectByIDStorage(ctx context.Context, id string) (model.Project, error) {
	collection := s.DB.Collection(projectCollection)

	project := model.Project{}

	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&project); err != nil {
		return model.Project{}, handleError("error occurred in FindOne", err)
	}

	return project, nil
}

func (s Storage) UpdateProjectByIDStorage(ctx context.Context, project model.Project) error {
	collection := s.DB.Collection(projectCollection)

	result, err := collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: project.ID}}, project)

	return handleUpdateByIDError(result, "error occurred in ReplaceOne", err)
}

func (s Storage) DeleteProjectByIDStorage(ctx context.Context, id string) error {
	collection := s.DB.Collection(projectCollection)

	result, err := collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})

	return handleDeleteByIDError(result, "error occurred in DeleteOne", err)
}
