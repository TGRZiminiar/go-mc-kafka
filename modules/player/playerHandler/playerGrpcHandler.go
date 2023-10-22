package playerhandler

import playerusecase "github.com/TGRZiminiar/go-mc-kafka/modules/player/playerUsecase"

type (
	playerGrpcHandlerService struct {
		playerUsecase playerusecase.PlayerUsecaseService
	}
)

func NewplayerGrpcHandler(playerUsecase playerusecase.PlayerUsecaseService) *playerGrpcHandlerService {
	return &playerGrpcHandlerService{playerUsecase}
}
