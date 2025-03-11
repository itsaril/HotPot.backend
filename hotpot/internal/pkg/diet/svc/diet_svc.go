package svc

import (
	"context"
	"log/slog"
)

type DietSvc struct {
	logger *slog.Logger
}

func NewDietService(logger *slog.Logger) *DietSvc {
	return &DietSvc{
		logger: logger,
	}
}

func (svc *DietSvc) Ping(_ context.Context) (bool, error) {
	return true, nil
}
