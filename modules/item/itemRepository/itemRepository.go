package itemRepository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/TGRZiminiar/go-mc-kafka/modules/item"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ItemRepositoryService interface {
		InsertOneItem(pctx context.Context, req *item.Item) (primitive.ObjectID, error)
		IsUniqueItem(pctx context.Context, title string) bool
		FindOneItem(pctx context.Context, itemId string) (*item.Item, error)
	}

	itemRepository struct {
		db *mongo.Client
	}
)

func NewItemRepository(db *mongo.Client) ItemRepositoryService {
	return &itemRepository{db}
}

func (r *itemRepository) itemDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("item_db")
}

func (r *itemRepository) IsUniqueItem(pctx context.Context, title string) bool {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.itemDbConn(ctx)
	col := db.Collection("items")

	result := new(item.Item)
	if err := col.FindOne(
		ctx,
		bson.M{"title": title},
	).Decode(result); err != nil {
		log.Printf("Error: IsUnique item : %s", err.Error())
		return true
	}

	return false
}
func (r *itemRepository) InsertOneItem(pctx context.Context, req *item.Item) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	db := r.itemDbConn(ctx)
	col := db.Collection("items")

	itemId, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOneItem: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one item fail")
	}

	return itemId.InsertedID.(primitive.ObjectID), nil

}

func (r *itemRepository) FindOneItem(pctx context.Context, itemId string) (*item.Item, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	db := r.itemDbConn(ctx)
	col := db.Collection("items")

	result := new(item.Item)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(itemId)}); err != nil {
		fmt.Printf("Error: Find One Item failed %s", err.Err())
		return nil, errors.New("error: item not found")
	}

	return result, nil

}
