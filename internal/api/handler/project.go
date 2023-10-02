package handler

import "golang.org/x/exp/slog"

type ProjectInterface interface {
}

type Project struct {
	Service ProjectInterface
	Log     *slog.Logger
}
