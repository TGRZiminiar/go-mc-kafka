package server

import (
	"log"

	inventoryhandler "github.com/TGRZiminiar/go-mc-kafka/modules/inventory/inventoryHandler"
	inventoryPb "github.com/TGRZiminiar/go-mc-kafka/modules/inventory/inventoryPb"
	inventoryrepository "github.com/TGRZiminiar/go-mc-kafka/modules/inventory/inventoryRepository"
	inventoryusecase "github.com/TGRZiminiar/go-mc-kafka/modules/inventory/inventoryUsecase"
	grpcconn "github.com/TGRZiminiar/go-mc-kafka/pkg/grpcConn"
)

func (s *server) inventoryService() {
	inventoryRepo := inventoryrepository.NewInventoryRepository(s.db)
	inventoryUsecase := inventoryusecase.NewInventoryUsecase(inventoryRepo)
	inventoryHttpHandler := inventoryhandler.NewInventoryHttpHandler(s.cfg, inventoryUsecase)
	inventoryGrpcHandler := inventoryhandler.NewInventoryGrpcHandler(inventoryUsecase)
	queueGrpcHandler := inventoryhandler.NewInventoryQueueHandler(s.cfg, inventoryUsecase)

	go queueGrpcHandler.AddPlayerItem()
	go queueGrpcHandler.RollbackAddPlayerItem()
	go queueGrpcHandler.RemovePlayerItem()
	go queueGrpcHandler.RollbackRemovePlayerItem()

	// gRPC
	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.InventoryUrl)

		inventoryPb.RegisterInventoryGrpcServiceServer(grpcServer, inventoryGrpcHandler)

		log.Printf("Inventory gRPC server listening on %s", s.cfg.Grpc.InventoryUrl)
		grpcServer.Serve(lis)
	}()

	_ = inventoryGrpcHandler
	_ = queueGrpcHandler

	inventory := s.app.Group("/inventory_v1")

	// HealthCheck

	inventory.GET("", s.healthCheckService)
	inventory.GET("/inventory/:player_id", inventoryHttpHandler.FindPlayerItems, s.middleware.JwtAuthorization, s.middleware.PlayerIdParamValidation)
}
