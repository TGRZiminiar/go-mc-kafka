package middlewarehandler

import (
	"github.com/TGRZiminiar/go-mc-kafka/config"
	middlewareUsecase "github.com/TGRZiminiar/go-mc-kafka/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface{}

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
