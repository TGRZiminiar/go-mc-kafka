package paymenthandler

import (
	"github.com/TGRZiminiar/go-mc-kafka/config"
	paymentusecase "github.com/TGRZiminiar/go-mc-kafka/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandlerService interface{}

	paymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase paymentusecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentusecase.PaymentUsecaseService) PaymentHttpHandlerService {
	return &paymentHttpHandler{cfg, paymentUsecase}
}
