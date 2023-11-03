package middlewarehandler

import (
	"net/http"
	"strings"

	"github.com/TGRZiminiar/go-mc-kafka/config"
	middlewareUsecase "github.com/TGRZiminiar/go-mc-kafka/modules/middleware/middlewareUsecase"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	MiddlewareHandlerService interface {
		JwtAuthorization(next echo.HandlerFunc) echo.HandlerFunc
		RbacAuthorization(next echo.HandlerFunc, expected []int) echo.HandlerFunc
		PlayerIdParamValidation(next echo.HandlerFunc) echo.HandlerFunc
	}

	MiddlewareHandler struct {
		cfg               *config.Config
		middlewareUsecase middlewareUsecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareUsecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &MiddlewareHandler{
		middlewareUsecase: middlewareUsecase,
		cfg:               cfg,
	}
}

func (h *MiddlewareHandler) JwtAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		accessToken := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")

		newCtx, err := h.middlewareUsecase.JwtAuthorization(c, h.cfg, accessToken)
		if err != nil {
			return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
		}
		return next(newCtx)
	}
}

func (h *MiddlewareHandler) RbacAuthorization(next echo.HandlerFunc, expected []int) echo.HandlerFunc {
	return func(c echo.Context) error {
		newCtx, err := h.middlewareUsecase.RbacAuthorization(c, h.cfg, expected)
		if err != nil {
			return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
		}
		return next(newCtx)
	}
}
func (h *MiddlewareHandler) PlayerIdParamValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		newCtx, err := h.middlewareUsecase.PlayerIdParamValidation(c)
		if err != nil {
			return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
		}
		return next(newCtx)
	}
}
