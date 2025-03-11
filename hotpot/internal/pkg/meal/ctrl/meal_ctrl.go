package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"hotpot/internal/core/utils/servers/http"
	"hotpot/internal/pkg/meal/svc"
	"log/slog"
)

type MealCtrl struct {
	logger  *slog.Logger
	mealSvc *svc.MealSvc
}

func NewMealController(logger *slog.Logger, svc *svc.MealSvc) *MealCtrl {
	return &MealCtrl{
		logger:  logger,
		mealSvc: svc,
	}
}

func (c *MealCtrl) Ping(ctx *fiber.Ctx) error {
	res, err := c.mealSvc.Ping(ctx.Context())
	if err != nil {
		return http.NewResponse(ctx, http.BadRequest, nil, http.CodeInternalError, "Something went wrong!")
	}
	return http.NewResponse(ctx, http.OK, res, 0, "")
}
