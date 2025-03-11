package svc

import (
	"context"
	"log/slog"
)

type UserSvc struct {
	logger *slog.Logger
}

func NewUserService(logger *slog.Logger) *UserSvc {
	return &UserSvc{
		logger: logger,
	}
}

func (svc *UserSvc) Ping(_ context.Context) (bool, error) {
	return true, nil
}
