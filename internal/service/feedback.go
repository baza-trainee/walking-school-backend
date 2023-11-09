package service

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/config"
	"github.com/baza-trainee/walking-school-backend/internal/model"
)

type Feedback struct {
	Cfg config.Feedback
}

func (f Feedback) CreateFeedbackService(ctx context.Context, fb model.Feedback) error {

	message := fmt.Sprintf(patternForFeedback, fb.Name, fb.Surname, fb.Phone, fb.Email, fb.Text)

	if err := sendMessage(f.Cfg.Host, f.Cfg.Port, f.Cfg.Username, f.Cfg.Password, fb.Email, f.Cfg.From, message); err != nil {
		return fmt.Errorf("error occurred in sendMessage: %w", err)
	}

	return nil
}
