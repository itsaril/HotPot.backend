package pkg

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"hotpot/internal/pkg/auth"
	"hotpot/internal/pkg/diet"
	"hotpot/internal/pkg/meal"
	"hotpot/internal/pkg/user"
	"log/slog"
)

type Module interface {
	InitHTTPRoutes(r fiber.Router)
}

type Router struct {
	logger  *slog.Logger
	modules []Module
}

func NewRouter(logger *slog.Logger) *Router {
	return &Router{
		logger: logger,
		modules: []Module{
			auth.New(logger),
			user.New(logger),
			diet.New(logger),
			meal.New(logger),
		},
	}
}

func (r *Router) RegisterModule(module Module) {
	r.modules = append(r.modules, module)
}

func (r *Router) Init(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	for _, module := range r.modules {
		module.InitHTTPRoutes(app)
	}
}
