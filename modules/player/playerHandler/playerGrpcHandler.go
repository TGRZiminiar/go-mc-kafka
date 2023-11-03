package playerhandler

import (
	"context"

	playerPb "github.com/TGRZiminiar/go-mc-kafka/modules/player/playerPb"
	playerusecase "github.com/TGRZiminiar/go-mc-kafka/modules/player/playerUsecase"
)

type (
	playerGrpcHandler struct {
		playerPb.UnimplementedPlayerGrpcServiceServer
		playerUsecase playerusecase.PlayerUsecaseService
	}
)

func NewplayerGrpcHandler(playerUsecase playerusecase.PlayerUsecaseService) *playerGrpcHandler {
	return &playerGrpcHandler{playerUsecase: playerUsecase}
}

func (g *playerGrpcHandler) CredentialSearch(ctx context.Context, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	return g.playerUsecase.FindOnePlayerCredentail(ctx, req.Password, req.Email)
}

func (g *playerGrpcHandler) FindOnePlayerProfileToRefresh(ctx context.Context, req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error) {
	return g.playerUsecase.FindOnePlayerProfileToRefresh(ctx, req.PlayerId)
}

func (g *playerGrpcHandler) GetPlayerSavingAccount(ctx context.Context, req *playerPb.GetPlayerSavingAccountReq) (*playerPb.GetPlayerSavingAccountRes, error) {
	return nil, nil

}
