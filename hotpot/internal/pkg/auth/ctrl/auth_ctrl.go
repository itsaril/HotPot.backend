package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"hotpot/internal/core/utils/servers/http"
	"hotpot/internal/pkg/auth/svc"
	"log/slog"
)

type AuthCtrl struct {
	logger  *slog.Logger
	authSvc *svc.AuthSvc
}

func NewAuthController(logger *slog.Logger, svc *svc.AuthSvc) *AuthCtrl {
	return &AuthCtrl{
		logger:  logger,
		authSvc: svc,
	}
}

func (c *AuthCtrl) Ping(ctx *fiber.Ctx) error {
	res, err := c.authSvc.Ping(ctx.Context())
	if err != nil {
		return http.NewResponse(ctx, http.BadRequest, nil, http.CodeInternalError, "Something went wrong!")
	}
	return http.NewResponse(ctx, http.OK, res, 0, "")
}
