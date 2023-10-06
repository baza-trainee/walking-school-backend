package handler

import (
	"errors"

	"github.com/baza-trainee/walking-school-backend/internal/model"
	"github.com/go-playground/validator"
)

var validate = validator.New() //nolint

const (
	standartLimitValue  = 10
	standartOffsetValue = 0
)

func UserValidate(validate *validator.Validate, user model.User) error {
	if user.Phone == "" && user.Email == "" {
		return errors.New("phone or email must be input")
	}

	if err := validate.Struct(user); err != nil {
		return err
	}

	return nil
}
