package authhandler

import (
	"context"

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

func (g *authGrpcHandler) AccessTokenSearch(ctx context.Context, req *authPb.AccessTokenSearchReq) (*authPb.AccessTokenSearchRes, error) {
	return nil, nil
}

func (g *authGrpcHandler) RolesCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return nil, nil
}
