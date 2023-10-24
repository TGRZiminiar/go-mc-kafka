package server

import (
	paymenthandler "github.com/TGRZiminiar/go-mc-kafka/modules/payment/paymentHandler"
	"github.com/TGRZiminiar/go-mc-kafka/modules/payment/paymentRepository"
	paymentusecase "github.com/TGRZiminiar/go-mc-kafka/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	paymentRepo := paymentRepository.NewPaymentRepository(s.db)
	paymentUsecase := paymentusecase.NewPaymentUsecase(paymentRepo)
	paymentHttpHandler := paymenthandler.NewPaymentHttpHandler(s.cfg, paymentUsecase)
	paymentQueueHandler := paymenthandler.NewPaymentQueueHandler(s.cfg, paymentUsecase)

	_ = paymentHttpHandler
	_ = paymentQueueHandler

	payment := s.app.Group("/payment_v1")

	// HealthCheck
	payment.GET("", s.healthCheckService)

}
