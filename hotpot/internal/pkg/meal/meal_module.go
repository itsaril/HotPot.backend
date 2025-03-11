package meal

import (
	"github.com/gofiber/fiber/v2"
	"hotpot/internal/pkg/meal/ctrl"
	"hotpot/internal/pkg/meal/svc"
	"log/slog"
)

type Module struct {
	Name    string
	Version string

	logger         *slog.Logger
	MealController *ctrl.MealCtrl
}

func New(logger *slog.Logger) *Module {
	mod := &Module{
		Name:    "meal-module",
		Version: "v1",
		logger:  logger,
		MealController: ctrl.NewMealController(
			logger,
			svc.NewMealService(logger),
		),
	}
	return mod
}

func (m *Module) InitHTTPRoutes(r fiber.Router) {
	root := r.Group("/" + m.Name).
		Group("/api").
		Group("/" + m.Version)

	modGroup := root.Group("/meal")
	modGroup.Get("/ping", m.MealController.Ping)
}
