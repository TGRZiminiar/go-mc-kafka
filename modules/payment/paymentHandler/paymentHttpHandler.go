package paymenthandler

import (
	"context"
	"net/http"

	"github.com/TGRZiminiar/go-mc-kafka/config"
	"github.com/TGRZiminiar/go-mc-kafka/modules/payment"
	paymentusecase "github.com/TGRZiminiar/go-mc-kafka/modules/payment/paymentUsecase"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/request"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	PaymentHttpHandlerService interface {
		BuyItem(c echo.Context) error
		SellItem(c echo.Context) error
	}

	paymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase paymentusecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentusecase.PaymentUsecaseService) PaymentHttpHandlerService {
	return &paymentHttpHandler{cfg, paymentUsecase}
}

func (h *paymentHttpHandler) BuyItem(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	playerId := c.Get("player_id").(string)

	req := &payment.ItemServiceReq{
		Items: make([]*payment.ItemServiceReqDatum, 0),
	}

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.paymentUsecase.BuyItem(ctx, h.cfg, req, playerId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
func (h *paymentHttpHandler) SellItem(c echo.Context) error {

	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	playerId := c.Get("player_id").(string)

	req := &payment.ItemServiceReq{
		Items: make([]*payment.ItemServiceReqDatum, 0),
	}

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.paymentUsecase.SellItem(ctx, h.cfg, req, playerId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
