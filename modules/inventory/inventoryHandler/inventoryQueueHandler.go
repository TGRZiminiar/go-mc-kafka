package inventoryhandler

import (
	"github.com/TGRZiminiar/go-mc-kafka/config"
	inventoryusecase "github.com/TGRZiminiar/go-mc-kafka/modules/inventory/inventoryUsecase"
)

type (
	InventoryQueueHandlerService interface{}

	inventoryQueueHandler struct {
		cfg              *config.Config
		inventoryUsecase inventoryusecase.InventoryUsecaseService
	}
)

func NewInventoryQueueHandler(cfg *config.Config, inventoryUsecase inventoryusecase.InventoryUsecaseService) InventoryQueueHandlerService {
	return &inventoryQueueHandler{cfg, inventoryUsecase}
}
