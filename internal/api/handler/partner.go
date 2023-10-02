package handler

import "golang.org/x/exp/slog"

type PartnerInterface interface {
}

type Partner struct {
	Service PartnerInterface
	Log     *slog.Logger
}
