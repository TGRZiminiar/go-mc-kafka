package paymentRepository

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/TGRZiminiar/go-mc-kafka/config"
	"github.com/TGRZiminiar/go-mc-kafka/modules/player"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/queue"
)

func (r *paymentRepository) DockedPlayerMoney(pctx context.Context, cfg *config.Config, req *player.CreatePlayerTransactionReq) error {

	reqInBytes, err := json.Marshal(req)
	if err != nil {
		log.Printf("Error: DockedPlayerMoney Failed %s", err.Error())
		return errors.New("error: docked player money failed")
	}

	if err := queue.PushMessageWithKeyToQueue(
		[]string{cfg.Kafka.Url},
		cfg.Kafka.ApiKey,
		cfg.Kafka.Secret,
		"player",
		"buy",
		reqInBytes,
	); err != nil {
		log.Printf("Error: DockedPlayerMoney Failed %s", err.Error())
		return errors.New("error: docked player money failed")
	}

	return nil
}

func (r *paymentRepository) RollBackDockedPlayerMoney(pctx context.Context, cfg *config.Config, req *player.CreatePlayerTransactionReq) error {

	reqInBytes, err := json.Marshal(req)
	if err != nil {
		log.Printf("Error: RollBackDockedPlayerMoney Failed %s", err.Error())
		return errors.New("error: roll back docked player money failed")
	}

	if err := queue.PushMessageWithKeyToQueue(
		[]string{cfg.Kafka.Url},
		cfg.Kafka.ApiKey,
		cfg.Kafka.Secret,
		"player",
		"rtransaction",
		reqInBytes,
	); err != nil {
		log.Printf("Error: RollBackDockedPlayerMoney Failed %s", err.Error())
		return errors.New("error: roll back docked player money failed")
	}

	return nil
}
