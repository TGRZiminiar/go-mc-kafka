package itemhandler

import (
	"context"

	itemPb "github.com/TGRZiminiar/go-mc-kafka/modules/item/itemPb"
	itemusecase "github.com/TGRZiminiar/go-mc-kafka/modules/item/itemUsecase"
)

type (
	itemGrpcHandler struct {
		itemUsecase itemusecase.ItemUsecaseService
		itemPb.UnimplementedItemGrpcServiceServer
	}
)

func NewItemGrpcHandler(itemUsecase itemusecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{itemUsecase: itemUsecase}
}

func (g *itemGrpcHandler) FindItemsInIds(ctx context.Context, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error) {
	return g.itemUsecase.FindItemInIds(ctx, req)
}
