package authhandler

import (
	"github.com/TGRZiminiar/go-mc-kafka/config"
	authusecase "github.com/TGRZiminiar/go-mc-kafka/modules/auth/authUsecase"
	"github.com/labstack/echo/v4"
)

type (
	AuthHttpHandlerService interface{}

	authHttpHandler struct {
		cfg         *config.Config
		authUseCase authusecase.AuthUseCaseService
	}
)

func NewAuthHttpHandler(pcfg *config.Config, authUseCase authusecase.AuthUseCaseService) AuthHttpHandlerService {
	return &authHttpHandler{
		authUseCase: authUseCase,
		cfg:         pcfg,
	}
}

func (h *authHttpHandler) Login(c echo.Context) error {

	// ctx := context.Background()

	// req := new(auth.PlayerLoginReq)

	return nil
}
