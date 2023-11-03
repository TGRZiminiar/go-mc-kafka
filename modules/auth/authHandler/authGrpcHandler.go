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
	return g.authUsecase.FindOneAccessToken(ctx, req.AccessToken)
}

func (g *authGrpcHandler) RolesCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {

	return g.authUsecase.RolesCount(ctx)
}
