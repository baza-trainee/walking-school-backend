package service

import (
	"context"
	"fmt"

	"github.com/baza-trainee/walking-school-backend/internal/config"
	"github.com/baza-trainee/walking-school-backend/internal/model"
	"gopkg.in/gomail.v2"
)

type Feedback struct {
	Cfg config.Form
}

func (f Feedback) CreateFormService(ctx context.Context, form model.Form) error {
	m := gomail.NewMessage()
	m.SetHeader("From", f.Cfg.Email)
	m.SetHeader("To", f.Cfg.Email)
	m.SetHeader("Subject", "Walking School Form")
	m.SetBody("text/html", fmt.Sprintf(
		`<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Пять строк текста</title>
		</head>
		<body>
			<p>Name: %s</p>
			<p>Surname: %s</p>
			<p>Email: %s</p>
			<p>Phone: %s</p>
			<p>Text: %s</p>
		</body>
		</html>`, form.Name, form.Surname, form.Email, form.Phone, form.Text))

	a := gomail.NewDialer("smtp.gmail.com", 587, f.Cfg.Email, f.Cfg.GeneratedCode)

	if err := a.DialAndSend(m); err != nil {
		return fmt.Errorf("error occurred during sending the form to email: %w", err)
	}

	return nil
}
