package server

import (
	"log"

	itemhandler "github.com/TGRZiminiar/go-mc-kafka/modules/item/itemHandler"
	itemPb "github.com/TGRZiminiar/go-mc-kafka/modules/item/itemPb"
	"github.com/TGRZiminiar/go-mc-kafka/modules/item/itemRepository"
	itemusecase "github.com/TGRZiminiar/go-mc-kafka/modules/item/itemUsecase"
	grpcconn "github.com/TGRZiminiar/go-mc-kafka/pkg/grpcConn"
)

func (s *server) itemService() {
	itemRepo := itemRepository.NewItemRepository(s.db)
	itemUsecase := itemusecase.NewItemUsecase(itemRepo)
	itemHttpHandler := itemhandler.NewItemHttpHandler(s.cfg, itemUsecase)
	itemGrpcHandler := itemhandler.NewItemGrpcHandler(itemUsecase)

	// gRPC
	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)

		itemPb.RegisterItemGrpcServiceServer(grpcServer, itemGrpcHandler)

		log.Printf("Item gRPC server listening on %s", s.cfg.Grpc.ItemUrl)
		grpcServer.Serve(lis)
	}()

	_ = itemGrpcHandler

	item := s.app.Group("/item_v1")

	// HealthCheck
	item.GET("", s.healthCheckService)
	item.POST("/create-item", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(itemHttpHandler.CreateItem, []int{1, 0})))
	item.GET("/item/:itemId", itemHttpHandler.FindOneItem)
}
