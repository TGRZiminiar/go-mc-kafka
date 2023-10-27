package authhandler

import (
	authPb "github.com/TGRZiminiar/go-mc-kafka/modules/auth/authPb"
	authusecase "github.com/TGRZiminiar/go-mc-kafka/modules/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		authUsecase authusecase.AuthUseCaseService
		authPb.UnimplementedAuthGrpcServiceServer
	}
)

func NewAuthGrpcHandler(authUsecase authusecase.AuthUseCaseService) *authGrpcHandler {
	return &authGrpcHandler{authUsecase: authUsecase}
}
