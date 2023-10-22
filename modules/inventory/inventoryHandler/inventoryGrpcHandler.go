package inventoryhandler

import inventoryusecase "github.com/TGRZiminiar/go-mc-kafka/modules/inventory/inventoryUsecase"

type (
	inventoryGrpcHandler struct {
		inventoryUsecase inventoryusecase.InventoryUsecaseService
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryusecase.InventoryUsecaseService) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{inventoryUsecase}
}
