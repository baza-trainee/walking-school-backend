package service

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/google/uuid"
)

type FeedbackStorageInterface interface {
	CreateFeedbackStorage(context.Context, model.Feedback) error
}

type Feedback struct {
	Storage FeedbackStorageInterface
}

func (f Feedback) CreateFeedbackService(ctx context.Context, feedback model.Feedback) error {
	feedback.ID = uuid.NewString()

	if err := f.Storage.CreateFeedbackStorage(ctx, feedback); err != nil {
		return fmt.Errorf("error occurred in CreateFeedbackStorage: %w", err)
	}

	// m := gomail.NewMessage()
	// m.SetHeader("From", "mi8aviation@gmail.com")
	// m.SetHeader("To", "etverya11@gmail.com")
	// m.SetHeader("Subject", "Title")
	// m.SetBody("text/html", feedback.Text)

	// a := gomail.NewDialer("smtp.gmail.com", 587, "mi8aviation@gmail.com", "bk321cu78ap")

	// if err := a.DialAndSend(m); err != nil {
	// 	return fmt.Errorf("error occurred during sending the message to email: %w", err)
	// }

	return nil
}
