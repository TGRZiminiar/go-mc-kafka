package playerhandler

import (
	"github.com/TGRZiminiar/go-mc-kafka/config"
	playerusecase "github.com/TGRZiminiar/go-mc-kafka/modules/player/playerUsecase"
)

type (
	PlayerQueueHandlerService interface{}

	playerQueueHandler struct {
		cfg           *config.Config
		playerUsecase playerusecase.PlayerUsecaseService
	}
)

func NewPlayerQueueHandler(cfg *config.Config, playerUsecase playerusecase.PlayerUsecaseService) PlayerQueueHandlerService {
	return &playerQueueHandler{cfg, playerUsecase}
}
