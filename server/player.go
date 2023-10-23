package server

import (
	playerhandler "github.com/TGRZiminiar/go-mc-kafka/modules/player/playerHandler"
	"github.com/TGRZiminiar/go-mc-kafka/modules/player/playerRepository"
	playerusecase "github.com/TGRZiminiar/go-mc-kafka/modules/player/playerUsecase"
)

func (s *server) playerService() {
	playerRepo := playerRepository.NewPlayerRepository(s.db)
	playerUsecase := playerusecase.NewPlayerUsecase(playerRepo)
	playerHttpHandler := playerhandler.NewPlayerHttpHandler(s.cfg, playerUsecase)
	playerGrpcHandler := playerhandler.NewplayerGrpcHandler(playerUsecase)
	queueGrpcHandler := playerhandler.NewPlayerQueueHandler(s.cfg, playerUsecase)

	_ = playerHttpHandler
	_ = playerGrpcHandler
	_ = queueGrpcHandler

	player := s.app.Group("/player_v1")

	// HealthCheck
	_ = player

}
