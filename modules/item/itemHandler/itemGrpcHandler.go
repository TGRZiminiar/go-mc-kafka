package itemhandler

import (
	itemusecase "github.com/TGRZiminiar/go-mc-kafka/modules/item/itemUsecase"
)

type (
	itemGrpcHandler struct {
		itemUsecase itemusecase.ItemUsecaseService
	}
)

func NewItemGrpcHandler(itemUsecase itemusecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{itemUsecase}
}
