package storage

import (
	"context"

	"github.com/baza-trainee/walking-school-backend/internal/model"
)

func (s Storage) CreateFeedbackStorage(ctx context.Context, feedback model.Feedback) error {
	collection := s.DB.Collection(feedbackCollection)

	_, err := collection.InsertOne(ctx, feedback)
	if err != nil {
		return handleError("error occurred in InsertOne", err)
	}

	return nil
}
