package server

import (
	itemhandler "github.com/TGRZiminiar/go-mc-kafka/modules/item/itemHandler"
	"github.com/TGRZiminiar/go-mc-kafka/modules/item/itemRepository"
	itemusecase "github.com/TGRZiminiar/go-mc-kafka/modules/item/itemUsecase"
)

func (s *server) itemService() {
	itemRepo := itemRepository.NewItemRepository(s.db)
	itemUsecase := itemusecase.NewItemUsecase(itemRepo)
	itemHttpHandler := itemhandler.NewItemHttpHandler(s.cfg, itemUsecase)
	itemGrpcHandler := itemhandler.NewItemGrpcHandler(itemUsecase)

	_ = itemHttpHandler
	_ = itemGrpcHandler

	item := s.app.Group("/item_v1")

	// HealthCheck
	item.GET("", s.healthCheckService)

}
