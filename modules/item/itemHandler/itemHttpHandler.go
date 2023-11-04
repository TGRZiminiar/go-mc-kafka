package itemhandler

import (
	"context"
	"net/http"

	"github.com/TGRZiminiar/go-mc-kafka/config"
	"github.com/TGRZiminiar/go-mc-kafka/modules/item"
	itemusecase "github.com/TGRZiminiar/go-mc-kafka/modules/item/itemUsecase"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/request"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	ItemHttpHandlerService interface {
		CreateItem(c echo.Context) error
		FindOneItem(c echo.Context) error
	}

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

func (h *itemHttpHandler) CreateItem(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	req := new(item.CreateItemReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	itemId, err := h.itemUsecase.CreateItem(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}
	return response.SuccessResponse(c, http.StatusCreated, itemId)
}

func (h *itemHttpHandler) FindOneItem(c echo.Context) error {
	ctx := context.Background()

	itemId := c.Param("itemId")

	res, err := h.itemUsecase.FindOneItem(ctx, itemId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}
	return response.SuccessResponse(c, http.StatusCreated, res)
}
