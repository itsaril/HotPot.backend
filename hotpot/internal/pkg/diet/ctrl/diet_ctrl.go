package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"hotpot/internal/core/utils/servers/http"
	"hotpot/internal/pkg/diet/svc"
	"log/slog"
)

type DietCtrl struct {
	logger  *slog.Logger
	dietSvc *svc.DietSvc
}

func NewDietController(logger *slog.Logger, svc *svc.DietSvc) *DietCtrl {
	return &DietCtrl{
		logger:  logger,
		dietSvc: svc,
	}
}

func (c *DietCtrl) Ping(ctx *fiber.Ctx) error {
	res, err := c.dietSvc.Ping(ctx.Context())
	if err != nil {
		return http.NewResponse(ctx, http.BadRequest, nil, http.CodeInternalError, "Something went wrong!")
	}
	return http.NewResponse(ctx, http.OK, res, 0, "")
}
