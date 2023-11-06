package paymentusecase

import (
	"context"
	"errors"
	"log"

	"github.com/IBM/sarama"
	"github.com/TGRZiminiar/go-mc-kafka/config"
	"github.com/TGRZiminiar/go-mc-kafka/modules/item"
	itemPb "github.com/TGRZiminiar/go-mc-kafka/modules/item/itemPb"
	"github.com/TGRZiminiar/go-mc-kafka/modules/payment"
	"github.com/TGRZiminiar/go-mc-kafka/modules/payment/paymentRepository"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/queue"
)

type (
	PaymentUsecaseService interface {
		GetOffset(pctx context.Context) (int64, error)
		UpsertOffset(pctx context.Context, offset int64) error
		FindItemsInIds(pctx context.Context, grpcUrl string, req []*payment.ItemServiceReqDatum) error
		PaymentConsumer(pctx context.Context, cfg *config.Config) (sarama.PartitionConsumer, error)
		BuyOrSellItemConsumer(pctx context.Context, key string, cfg *config.Config, resCh chan<- *payment.PaymentTransferRes)
	}

	paymentUsecase struct {
		paymentRepository paymentRepository.PaymentRepositoryService
	}
)

func NewPaymentUsecase(paymentRepository paymentRepository.PaymentRepositoryService) PaymentUsecaseService {
	return &paymentUsecase{paymentRepository}
}

func (u *paymentUsecase) FindItemsInIds(pctx context.Context, grpcUrl string, req []*payment.ItemServiceReqDatum) error {

	setIds := make(map[string]bool)
	for _, v := range req {
		if !setIds[v.ItemId] {
			setIds[v.ItemId] = true
		}
	}

	itemData, err := u.paymentRepository.FindItemInIds(pctx, grpcUrl, &itemPb.FindItemsInIdsReq{
		Ids: func() []string {
			itemIds := make([]string, 0)
			for k := range setIds {
				itemIds = append(itemIds, k)
			}
			return itemIds
		}(),
	})

	if err != nil {
		log.Printf("Error: Find Item In Ids Failed %s", err.Error())
		return err
	}

	itemMaps := make(map[string]*item.ItemShowCase)
	for _, v := range itemData.Items {
		itemMaps[v.Id] = &item.ItemShowCase{
			ItemId:   v.Id,
			Title:    v.Title,
			Price:    v.Price,
			ImageUrl: v.ImageUrl,
			Damage:   int(v.Damage),
		}
	}

	for i := range req {
		if _, ok := itemMaps[req[i].ItemId]; !ok {
			log.Printf("Error: Find Item In Ids failed No Item Id Found In Result")
			return errors.New("error: no item id in resukt")
		}
		req[i].Price = itemMaps[req[i].ItemId].Price
	}

	return nil
}

// Receive Message From Other Producer
func (u *paymentUsecase) PaymentConsumer(pctx context.Context, cfg *config.Config) (sarama.PartitionConsumer, error) {

	worker, err := queue.ConnectConsumer([]string{cfg.Kafka.Url}, cfg.Kafka.ApiKey, cfg.Kafka.Secret)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error: payment consumer connection failed")
	}

	offset, err := u.paymentRepository.GetOffset(pctx)
	if err != nil {
		return nil, err
	}

	consumer, err := worker.ConsumePartition("payment", 0, offset)
	if err != nil {
		log.Println("Trying to set offset to 0")
		consumer, err = worker.ConsumePartition("payment", 0, 0)
		if err != nil {
			log.Printf("Error: Payment Consumer Failed %s", err.Error())
			return nil, err
		}
	}

	return consumer, nil
}

// Consumer Wait For Connection From Other Microservice
func (u *paymentUsecase) BuyOrSellItemConsumer(pctx context.Context, key string, cfg *config.Config, resCh chan<- *payment.PaymentTransferRes) {
	consumer, err := u.PaymentConsumer(pctx, cfg)
	if err != nil {
		resCh <- nil
		return
	}

	log.Println("Start BuyOrSellItemConsumer ...")

	select {
	case err := <-consumer.Errors():
		log.Printf("Error: BuyOrSellItemConsumer failed %s", err.Error())
		resCh <- nil
		return
	case msg := <-consumer.Messages():
		if string(msg.Key) == key {
			u.UpsertOffset(pctx, msg.Offset+1)

			req := new(payment.PaymentTransferRes)

			if err := queue.DecodeMessage(req, msg.Value); err != nil {
				resCh <- nil
				return
			}
			resCh <- req
			log.Printf("BuyItemConsumer | Topic(%s) | Offset(%d) | Message(%s) \n", msg.Topic, msg.Offset, msg.Value)
		}

	}

}

func (u *paymentUsecase) BuyItem(pctx context.Context, cfg *config.Config, req *payment.ItemServiceReq, playerId string) (*payment.PaymentTransferRes, error) {

	if err := u.FindItemsInIds(pctx, cfg.Grpc.ItemUrl, req.Items); err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *paymentUsecase) SellItem(pctx context.Context, cfg *config.Config, req *payment.ItemServiceReq, playerId string) (*payment.PaymentTransferRes, error) {

	if err := u.FindItemsInIds(pctx, cfg.Grpc.ItemUrl, req.Items); err != nil {
		return nil, err
	}

	return nil, nil
}

// Kafka Func
func (u *paymentUsecase) GetOffset(pctx context.Context) (int64, error) {
	return u.paymentRepository.GetOffset(pctx)
}
func (u *paymentUsecase) UpsertOffset(pctx context.Context, offset int64) error {
	return u.paymentRepository.UpsertOffset(pctx, offset)
}
