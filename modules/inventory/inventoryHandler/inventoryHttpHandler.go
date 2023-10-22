package inventoryhandler

import (
	"github.com/TGRZiminiar/go-mc-kafka/config"
	inventoryusecase "github.com/TGRZiminiar/go-mc-kafka/modules/inventory/inventoryUsecase"
)

type (
	InventoryHttpHandlerService interface{}

	inventoryHttpHandler struct {
		cfg              *config.Config
		inventoryUsecase inventoryusecase.InventoryUsecaseService
	}
)

func NewInventoryHttpHandler(cfg *config.Config, inventoryUsecase inventoryusecase.InventoryUsecaseService) InventoryHttpHandlerService {
	return &inventoryHttpHandler{cfg, inventoryUsecase}
}
