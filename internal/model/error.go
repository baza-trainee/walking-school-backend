package model

import "errors"

var (
	ErrNotFound       = errors.New("nothing was found")
	ErrNoContent      = errors.New("no content")
	ErrNothinChanged  = errors.New("nothing changed")
	ErrRequestTimeout = errors.New("request timeout")
	ErrBadRequest     = errors.New("bad request")
	ErrConflict       = errors.New("conflict")
)
