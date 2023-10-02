package handler

import "golang.org/x/exp/slog"

type UserInterface interface {
}

type User struct {
	Service UserInterface
	Log     *slog.Logger
}
