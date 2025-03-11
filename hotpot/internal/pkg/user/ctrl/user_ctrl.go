package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"hotpot/internal/core/utils/servers/http"
	"hotpot/internal/pkg/user/svc"
	"log/slog"
)

type UserCtrl struct {
	logger  *slog.Logger
	userSvc *svc.UserSvc
}

func NewUserController(logger *slog.Logger, svc *svc.UserSvc) *UserCtrl {
	return &UserCtrl{
		logger:  logger,
		userSvc: svc,
	}
}

func (c *UserCtrl) Ping(ctx *fiber.Ctx) error {
	res, err := c.userSvc.Ping(ctx.Context())
	if err != nil {
		return http.NewResponse(ctx, http.BadRequest, nil, http.CodeInternalError, "Something went wrong!")
	}
	return http.NewResponse(ctx, http.OK, res, 0, "")
}
