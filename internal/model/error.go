package model

import "errors"

var (
	ErrNotFound             = errors.New("nothing was found")
	ErrNothinChanged        = errors.New("nothing changed")
	ErrRequestTimeout       = errors.New("request timeout")
	ErrBadRequest           = errors.New("bad request")
	ErrConflict             = errors.New("conflict")
	ErrInvalidSigningMethod = errors.New("invalid signing metod")
	ErrWrongTokenClaimType  = errors.New("wrong token claims")
)
