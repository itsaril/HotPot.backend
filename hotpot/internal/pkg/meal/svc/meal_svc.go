package svc

import (
	"context"
	"log/slog"
)

type MealSvc struct {
	logger *slog.Logger
}

func NewMealService(logger *slog.Logger) *MealSvc {
	return &MealSvc{
		logger: logger,
	}
}

func (svc *MealSvc) Ping(_ context.Context) (bool, error) {
	return true, nil
}
