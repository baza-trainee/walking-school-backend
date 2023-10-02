package handler

import "golang.org/x/exp/slog"

type ServiceInterface interface {
}

type Handler struct {
	Project
	User
	Partner
}

func NewHandler(service ServiceInterface, log *slog.Logger) (Handler, error) {
	return Handler{
		Project{Service: service, Log: log},
		User{Service: service, Log: log},
		Partner{Service: service, Log: log},
	}, nil
}
