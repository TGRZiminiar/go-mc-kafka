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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	ItemRepositoryService interface {
		InsertOneItem(pctx context.Context, req *item.Item) (primitive.ObjectID, error)
		IsUniqueItem(pctx context.Context, title string) bool
		FindOneItem(pctx context.Context, itemId string) (*item.Item, error)
		FindManyItem(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*item.ItemShowCase, error)
		CountItems(pctx context.Context, filter primitive.D) (int64, error)
		UpdateOneItem(pctx context.Context, itemId string, req primitive.M) error
		EnableOrDisableItem(pctx context.Context, itemId string, isActive bool) error
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
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(itemId)}).Decode(result); err != nil {
		fmt.Printf("Error: Find One Item failed %s", err.Error())
		return nil, errors.New("error: item not found")
	}

	return result, nil

}

func (r *itemRepository) FindManyItem(pctx context.Context, filter primitive.D, opts []*options.FindOptions) ([]*item.ItemShowCase, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	db := r.itemDbConn(ctx)
	col := db.Collection("items")

	cursors, err := col.Find(ctx, filter, opts...)
	if err != nil {
		log.Printf("Error: Find Many Item Failed %s", err.Error())
		return make([]*item.ItemShowCase, 0), errors.New("error: find many item failed")
	}

	results := make([]*item.ItemShowCase, 0)
	for cursors.Next(ctx) {
		result := new(item.Item)
		if err := cursors.Decode(result); err != nil {
			log.Printf("Error: Find Many Item Failed %s", err.Error())
			return make([]*item.ItemShowCase, 0), errors.New("error: find many item failed")
		}
		results = append(results, &item.ItemShowCase{
			ItemId:   "item:" + result.Id.Hex(),
			Title:    result.Title,
			Price:    result.Price,
			Damage:   result.Damage,
			ImageUrl: result.ImageUrl,
		})
	}
	return results, nil

}
func (r *itemRepository) CountItems(pctx context.Context, filter primitive.D) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	db := r.itemDbConn(ctx)
	col := db.Collection("items")

	count, err := col.CountDocuments(ctx, filter)
	if err != nil {
		return -1, errors.New("error: count item failed")
	}
	return count, nil

}

func (r *itemRepository) UpdateOneItem(pctx context.Context, itemId string, req primitive.M) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	db := r.itemDbConn(ctx)
	col := db.Collection("items")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(itemId)}, bson.M{"$set": req})
	if err != nil {
		log.Printf("Error: Update One Item Failed %s", err.Error())
		return errors.New("error: update one item failed")
	}
	log.Println("Update One Item Success", result)
	return nil
}

func (r *itemRepository) EnableOrDisableItem(pctx context.Context, itemId string, isActive bool) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.itemDbConn(ctx)
	col := db.Collection("items")

	result, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertToObjectId(itemId)}, bson.M{"$set": bson.M{"usage_status": isActive}})
	if err != nil {
		log.Printf("Error: EnableOrDisableItem failed: %s", err.Error())
		return errors.New("error: enable or disable item failed")
	}
	log.Printf("EnableOrDisableItem result: %v", result.ModifiedCount)

	return nil
}
