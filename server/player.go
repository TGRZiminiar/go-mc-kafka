package server

import (
	"log"

	playerhandler "github.com/TGRZiminiar/go-mc-kafka/modules/player/playerHandler"
	playerPb "github.com/TGRZiminiar/go-mc-kafka/modules/player/playerPb"
	"github.com/TGRZiminiar/go-mc-kafka/modules/player/playerRepository"
	playerusecase "github.com/TGRZiminiar/go-mc-kafka/modules/player/playerUsecase"
	grpcconn "github.com/TGRZiminiar/go-mc-kafka/pkg/grpcConn"
)

func (s *server) playerService() {
	playerRepo := playerRepository.NewPlayerRepository(s.db)
	playerUsecase := playerusecase.NewPlayerUsecase(playerRepo)
	playerHttpHandler := playerhandler.NewPlayerHttpHandler(s.cfg, playerUsecase)
	playerGrpcHandler := playerhandler.NewplayerGrpcHandler(playerUsecase)
	queueGrpcHandler := playerhandler.NewPlayerQueueHandler(s.cfg, playerUsecase)

	// gRPC
	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.PlayerUrl)

		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, playerGrpcHandler)

		log.Printf("Player gRPC server listening on %s", s.cfg.Grpc.PlayerUrl)
		grpcServer.Serve(lis)
	}()

	_ = playerHttpHandler
	_ = playerGrpcHandler
	_ = queueGrpcHandler

	player := s.app.Group("/player_v1")

	// HealthCheck
	player.GET("", s.healthCheckService)

	player.POST("/player/register", playerHttpHandler.CreatePlayer)
	player.GET("/player/:player_id", playerHttpHandler.FindOnePlayerProfile)
	player.POST("/player/add-money", playerHttpHandler.AddPlayerMoney)

}
