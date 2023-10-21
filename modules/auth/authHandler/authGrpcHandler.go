package authhandler

import authusecase "github.com/TGRZiminiar/go-mc-kafka/modules/auth/authUsecase"

type (
	authGrpcHandler struct {
		authUsecase authusecase.AuthUseCaseService
	}
)

func NewAuthGrpcHandler(authUsecase authusecase.AuthUseCaseService) authusecase.AuthUseCaseService {
	return &authGrpcHandler{authUsecase}
}
