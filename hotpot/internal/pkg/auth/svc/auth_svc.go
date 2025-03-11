package svc

import (
	"context"
	"log/slog"
)

type AuthSvc struct {
	logger *slog.Logger
}

func NewAuthService(logger *slog.Logger) *AuthSvc {
	return &AuthSvc{
		logger: logger,
	}
}

func (svc *AuthSvc) Ping(_ context.Context) (bool, error) {
	return true, nil
}
