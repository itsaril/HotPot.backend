package auth

import (
	"github.com/gofiber/fiber/v2"
	"hotpot/internal/pkg/auth/ctrl"
	"hotpot/internal/pkg/auth/svc"
	"log/slog"
)

type Module struct {
	Name    string
	Version string

	logger         *slog.Logger
	AuthController *ctrl.AuthCtrl
}

func New(logger *slog.Logger) *Module {
	mod := &Module{
		Name:    "auth-module",
		Version: "v1",
		logger:  logger,
		AuthController: ctrl.NewAuthController(
			logger,
			svc.NewAuthService(logger),
		),
	}
	return mod
}

func (m *Module) InitHTTPRoutes(r fiber.Router) {
	root := r.Group("/" + m.Name).
		Group("/api").
		Group("/" + m.Version)

	modGroup := root.Group("/auth")
	modGroup.Get("/ping", m.AuthController.Ping)
}
