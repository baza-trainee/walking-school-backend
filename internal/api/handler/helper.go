package handler

import (
	"github.com/go-playground/validator"
)

var validate = validator.New() //nolint

const (
	standartLimitValue  = 10
	standartOffsetValue = 0
)
