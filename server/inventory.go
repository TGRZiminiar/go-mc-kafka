package server

import (
	inventoryhandler "github.com/TGRZiminiar/go-mc-kafka/modules/inventory/inventoryHandler"
	inventoryrepository "github.com/TGRZiminiar/go-mc-kafka/modules/inventory/inventoryRepository"
	inventoryusecase "github.com/TGRZiminiar/go-mc-kafka/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	inventoryRepo := inventoryrepository.NewInventoryRepository(s.db)
	inventoryUsecase := inventoryusecase.NewInventoryUsecase(inventoryRepo)
	inventoryHttpHandler := inventoryhandler.NewInventoryHttpHandler(s.cfg, inventoryUsecase)
	inventoryGrpcHandler := inventoryhandler.NewInventoryGrpcHandler(inventoryUsecase)
	queueGrpcHandler := inventoryhandler.NewInventoryQueueHandler(s.cfg, inventoryUsecase)

	_ = inventoryHttpHandler
	_ = inventoryGrpcHandler
	_ = queueGrpcHandler

	inventory := s.app.Group("/inventory_v1")

	// HealthCheck
	_ = inventory

}
