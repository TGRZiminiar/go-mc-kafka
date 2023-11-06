package playerRepository

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/TGRZiminiar/go-mc-kafka/config"
	"github.com/TGRZiminiar/go-mc-kafka/modules/payment"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/queue"
)

// Producer that produce message to kafka
// This func will produce to BuyOrSellConsumer Func in paymentusecase.go
func (r *playerRepository) DockedPlayerMoneyRes(pctx context.Context, cfg *config.Config, req *payment.PaymentTransferRes) error {
	reqInBytes, err := json.Marshal(req)
	if err != nil {
		log.Printf("Error: DockedPlayerMoneyRes Failed %s", err.Error())
		return errors.New("error: docked player money res failed")
	}

	if err := queue.PushMessageWithKeyToQueue(
		[]string{cfg.Kafka.Url},
		cfg.Kafka.ApiKey,
		cfg.Kafka.Secret,
		"payment",
		"buy",
		reqInBytes,
	); err != nil {
		log.Printf("Error: DockedPlayerMoneyRes Failed %s", err.Error())
		return errors.New("error: docked player money res failed")
	}

	return nil
}
