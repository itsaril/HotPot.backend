package diet

import (
	"github.com/gofiber/fiber/v2"
	"hotpot/internal/pkg/diet/ctrl"
	"hotpot/internal/pkg/diet/svc"
	"log/slog"
)

type Module struct {
	Name    string
	Version string

	logger         *slog.Logger
	DietController *ctrl.DietCtrl
}

func New(logger *slog.Logger) *Module {
	mod := &Module{
		Name:    "diet-module",
		Version: "v1",
		logger:  logger,
		DietController: ctrl.NewDietController(
			logger,
			svc.NewDietService(logger),
		),
	}
	return mod
}

func (m *Module) InitHTTPRoutes(r fiber.Router) {
	root := r.Group("/" + m.Name).
		Group("/api").
		Group("/" + m.Version)

	modGroup := root.Group("/diet")
	modGroup.Get("/ping", m.DietController.Ping)
}
