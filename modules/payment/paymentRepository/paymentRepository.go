package paymentRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/TGRZiminiar/go-mc-kafka/config"
	itemPb "github.com/TGRZiminiar/go-mc-kafka/modules/item/itemPb"
	"github.com/TGRZiminiar/go-mc-kafka/modules/models"
	"github.com/TGRZiminiar/go-mc-kafka/modules/player"

	grpcconn "github.com/TGRZiminiar/go-mc-kafka/pkg/grpcConn"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/jwtauth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	PaymentRepositoryService interface {
		GetOffset(pctx context.Context) (int64, error)
		UpsertOffset(pctx context.Context, offset int64) error
		FindItemInIds(pctx context.Context, grpcUrl string, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error)
		DockedPlayerMoney(pctx context.Context, cfg *config.Config, req *player.CreatePlayerTransactionReq) error
		RollBackDockedPlayerMoney(pctx context.Context, cfg *config.Config, req *player.CreatePlayerTransactionReq) error
	}

	paymentRepository struct {
		db *mongo.Client
	}
)

func NewPaymentRepository(db *mongo.Client) PaymentRepositoryService {
	return &paymentRepository{db}
}

func (r *paymentRepository) paymentDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("payment_db")
}

func (r *paymentRepository) FindItemInIds(pctx context.Context, grpcUrl string, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	jwtauth.SetApiKeyInContext(&ctx)
	conn, err := grpcconn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed %s", err.Error())
		return nil, errors.New("error: grpc connection failed")
	}

	result, err := conn.Item().FindItemsInIds(ctx, req)
	if err != nil {
		log.Printf("Error: Find Item In Ids Failed %s", err.Error())
		return nil, errors.New("error: item not found")
	}

	if result == nil {
		log.Printf("Error: Find Item In Ids Failed %s", err.Error())
		return nil, errors.New("error: item not found")
	}

	if len(result.Items) == 0 {
		log.Printf("Error: Find Item In Ids Failed")
		return nil, errors.New("error: item not found")
	}

	return result, nil

}

// Kafka Func
func (r *paymentRepository) GetOffset(pctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	db := r.paymentDbConn(ctx)
	col := db.Collection("payment_queue")

	result := new(models.KafkaOffset)
	if err := col.FindOne(ctx, bson.M{}).Decode(result); err != nil {
		log.Printf("Error: Get Off Set Failed %s", err.Error())
		return -1, errors.New("error: get off set failed")
	}

	return result.Offset, nil
}

func (r *paymentRepository) UpsertOffset(pctx context.Context, offset int64) error {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	db := r.paymentDbConn(ctx)
	col := db.Collection("payment_queue")

	// option setUpsert true = if no data it will get insert instead
	result, err := col.UpdateOne(ctx, bson.M{}, bson.M{"$set": bson.M{"offset": offset}}, options.Update().SetUpsert(true))
	if err != nil {
		log.Printf("Error: Upsert OffSet Failed %s", err.Error())
		return errors.New("error: upsert offset failed")
	}
	log.Printf("Info: Upsert Offset Success: %v", result)

	return nil
}
