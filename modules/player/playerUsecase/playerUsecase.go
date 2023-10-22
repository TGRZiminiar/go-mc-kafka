package playerusecase

import (
	"github.com/TGRZiminiar/go-mc-kafka/modules/player/playerRepository"
)

type (
	PlayerUsecaseService interface{}

	playerUsecase struct {
		playerRepository playerRepository.PlayerRepositoryService
	}
)

func NewPlayerUsecase(playerRepository playerRepository.PlayerRepositoryService) PlayerUsecaseService {
	return &playerUsecase{playerRepository}
}
