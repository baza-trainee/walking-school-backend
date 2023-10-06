package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s Storage) CreateProjectStorage(ctx context.Context, project model.Project) error {
	collection := s.DB.Collection(projectCollection)

	_, err := collection.InsertOne(ctx, project)
	if err != nil {
		// Какие ошибки могут возвращаться?
		return fmt.Errorf("error occurred in InsertOne: %w", err)
	}

	return nil
}

func (s Storage) GetAllProjectStorage(ctx context.Context, query model.ProjectQuery) ([]model.Project, error) {
	collection := s.DB.Collection(projectCollection)

	result := make([]model.Project, 0)

	cursor, err := collection.Find(ctx, bson.D{}, LimitAndOffset(query.Limit, query.Offset))
	if err != nil {
		// Какие ошибки могут возвращаться?
		return nil, fmt.Errorf("error occurred in Find: %w", err)
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
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Project{}, model.ErrNotFound
		}

		return model.Project{}, fmt.Errorf("error occurred in FindOne: %w", err)
	}

	return project, nil
}

func (s Storage) UpdateProjectByIDStorage(ctx context.Context, project model.Project) error {
	collection := s.DB.Collection(projectCollection)

	result, err := collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: project.ID}}, project)
	if err != nil {
		return fmt.Errorf("error occurred in ReplaceOne: %w", err)
	}

	if result.MatchedCount != matchedOneDocument {
		return model.ErrNotFound
	}

	return nil
}

func (s Storage) DeleteProjectByIDStorage(ctx context.Context, id string) error {
	collection := s.DB.Collection(projectCollection)

	result, err := collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return fmt.Errorf("error occurred in DeleteOne: %w", err)
	}

	if result.DeletedCount != matchedOneDocument {
		return model.ErrNotFound
	}

	return nil
}
