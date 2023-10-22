package itemhandler

import (
	"github.com/TGRZiminiar/go-mc-kafka/config"
	itemusecase "github.com/TGRZiminiar/go-mc-kafka/modules/item/itemUsecase"
)

type (
	ItemHttpHandlerService interface{}

	itemHttpHandler struct {
		cfg         *config.Config
		itemUsecase itemusecase.ItemUsecaseService
	}
)

func NewItemHttpHandler(cfg *config.Config, itemUsecase itemusecase.ItemUsecaseService) ItemHttpHandlerService {
	return &itemHttpHandler{
		cfg,
		itemUsecase,
	}
}
