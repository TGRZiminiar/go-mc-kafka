package paymenthandler

import (
	"github.com/TGRZiminiar/go-mc-kafka/config"
	paymentusecase "github.com/TGRZiminiar/go-mc-kafka/modules/payment/paymentUsecase"
)

type (
	PaymentQueueHandlerService interface{}

	paymentQueueHandler struct {
		cfg            *config.Config
		paymentUsecase paymentusecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandler(cfg *config.Config, paymentUsecase paymentusecase.PaymentUsecaseService) PaymentQueueHandlerService {
	return &paymentQueueHandler{cfg, paymentUsecase}
}
