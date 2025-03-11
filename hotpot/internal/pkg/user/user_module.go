package user

import (
	"github.com/gofiber/fiber/v2"
	"hotpot/internal/pkg/user/ctrl"
	"hotpot/internal/pkg/user/svc"
	"log/slog"
)

type Module struct {
	Name    string
	Version string

	logger         *slog.Logger
	UserController *ctrl.UserCtrl
}

func New(logger *slog.Logger) *Module {
	mod := &Module{
		Name:    "user-module",
		Version: "v1",
		logger:  logger,
		UserController: ctrl.NewUserController(
			logger,
			svc.NewUserService(logger),
		),
	}
	return mod
}

func (m *Module) InitHTTPRoutes(r fiber.Router) {
	root := r.Group("/" + m.Name).
		Group("/api").
		Group("/" + m.Version)

	modGroup := root.Group("/user")
	modGroup.Get("/ping", m.UserController.Ping)
}
