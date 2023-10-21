package authhandler

import (
	"github.com/TGRZiminiar/go-mc-kafka/config"
	authusecase "github.com/TGRZiminiar/go-mc-kafka/modules/auth/authUsecase"
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
