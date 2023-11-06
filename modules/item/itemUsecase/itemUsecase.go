package itemusecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/TGRZiminiar/go-mc-kafka/modules/item"
	itemPb "github.com/TGRZiminiar/go-mc-kafka/modules/item/itemPb"
	"github.com/TGRZiminiar/go-mc-kafka/modules/item/itemRepository"
	"github.com/TGRZiminiar/go-mc-kafka/modules/models"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	ItemUsecaseService interface {
		CreateItem(pctx context.Context, req *item.CreateItemReq) (any, error)
		FindOneItem(pctx context.Context, itemId string) (*item.ItemShowCase, error)
		FindManyItem(pctx context.Context, basePaginateUrl string, req *item.ItemSearchReq) (*models.PaginateRes, error)
		EditItem(pctx context.Context, itemId string, req *item.ItemUpdateReq) (*item.ItemShowCase, error)
		EnableOrDisableItem(pctx context.Context, itemId string) (bool, error)
		FindItemInIds(pctx context.Context, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error)
	}

	itemUsecase struct {
		itemRepository itemRepository.ItemRepositoryService
	}
)

func NewItemUsecase(itemRepository itemRepository.ItemRepositoryService) ItemUsecaseService {
	return &itemUsecase{itemRepository}
}

func (u *itemUsecase) CreateItem(pctx context.Context, req *item.CreateItemReq) (any, error) {

	if !u.itemRepository.IsUniqueItem(pctx, req.Title) {
		return nil, errors.New("error; title already exist")
	}

	itemId, err := u.itemRepository.InsertOneItem(pctx, &item.Item{
		Title:       req.Title,
		Price:       req.Price,
		Damage:      req.Damage,
		ImageUrl:    req.ImageUrl,
		UsageStatus: true,
		CreatedAt:   utils.LocalTime(),
		UpdatedAt:   utils.LocalTime(),
	})
	if err != nil {
		return nil, err
	}
	return itemId.Hex(), nil
}

func (u *itemUsecase) FindOneItem(pctx context.Context, itemId string) (*item.ItemShowCase, error) {
	result, err := u.itemRepository.FindOneItem(pctx, itemId)
	if err != nil {
		return nil, err
	}
	return &item.ItemShowCase{
		ItemId:   "item:" + result.Id.Hex(),
		Title:    result.Title,
		Price:    result.Price,
		Damage:   result.Damage,
		ImageUrl: result.ImageUrl,
	}, nil
}

func (u *itemUsecase) FindManyItem(pctx context.Context, basePaginateUrl string, req *item.ItemSearchReq) (*models.PaginateRes, error) {

	findItemsFilter := bson.D{}
	findItemsOpts := make([]*options.FindOptions, 0)

	countItemsFilter := bson.D{}

	// Find Many Item
	if req.Start != "" {
		req.Start = strings.TrimPrefix(req.Start, "item:")
		findItemsFilter = append(findItemsFilter, bson.E{"_id", bson.D{{"$gt", utils.ConvertToObjectId(req.Start)}}})
	}
	if req.Title != "" {
		// Options: "i" = sensitive case
		findItemsFilter = append(findItemsFilter, bson.E{"title", primitive.Regex{Pattern: req.Title, Options: "i"}})
		countItemsFilter = append(countItemsFilter, bson.E{"title", primitive.Regex{Pattern: req.Title, Options: "i"}})
	}

	findItemsFilter = append(findItemsFilter, bson.E{"usage_status", true})
	countItemsFilter = append(countItemsFilter, bson.E{"usage_status", true})

	// Options
	findItemsOpts = append(findItemsOpts, options.Find().SetSort(bson.D{{"_id", 1}}))
	findItemsOpts = append(findItemsOpts, options.Find().SetLimit(int64(req.Limit)))

	// Find Item
	result, err := u.itemRepository.FindManyItem(pctx, findItemsFilter, findItemsOpts)
	if err != nil {
		return nil, err
	}

	// Count Item
	count, err := u.itemRepository.CountItems(pctx, countItemsFilter)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return &models.PaginateRes{
			Data:  make([]*item.ItemShowCase, 0),
			Total: count,
			First: models.FirstPaginate{
				Href: fmt.Sprintf("%s?title=%s&limit=%d", basePaginateUrl, req.Title, req.Limit),
			},
			Next: models.NextPaginate{
				Start: "",
				Href:  "",
			},
		}, nil
	}

	return &models.PaginateRes{
		Data:  result,
		Total: count,
		First: models.FirstPaginate{
			Href: fmt.Sprintf("%s?title=%s&limit=%d", basePaginateUrl, req.Title, req.Limit),
		},
		Limit: req.Limit,
		Next: models.NextPaginate{
			Start: result[len(result)-1].ItemId,
			Href:  fmt.Sprintf("%s?title=%s&start=%s&limit=%d", basePaginateUrl, req.Title, result[len(result)-1].ItemId, req.Limit),
		},
	}, nil

}

func (u *itemUsecase) EditItem(pctx context.Context, itemId string, req *item.ItemUpdateReq) (*item.ItemShowCase, error) {

	updateReq := bson.M{}

	if req.Title != "" {
		if !u.itemRepository.IsUniqueItem(pctx, req.Title) {
			log.Println("Error: This title already exist")
			return nil, errors.New("error: this title is already exist")
		}
		updateReq["title"] = req.Title
	}
	if req.ImageUrl != "" {
		updateReq["image_url"] = req.ImageUrl
	}
	if req.Damage > 0 {
		updateReq["damage"] = req.Damage
	}
	if req.Price >= 0 {
		updateReq["price"] = req.Price
	}

	updateReq["updated_at"] = utils.LocalTime()

	if err := u.itemRepository.UpdateOneItem(pctx, itemId, updateReq); err != nil {
		return nil, err
	}

	return u.FindOneItem(pctx, itemId)
}

func (u *itemUsecase) EnableOrDisableItem(pctx context.Context, itemId string) (bool, error) {
	result, err := u.itemRepository.FindOneItem(pctx, itemId)
	if err != nil {
		return false, err
	}

	if err := u.itemRepository.EnableOrDisableItem(pctx, itemId, !result.UsageStatus); err != nil {
		return false, err
	}

	return !result.UsageStatus, nil
}

func (u *itemUsecase) FindItemInIds(pctx context.Context, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error) {
	filter := bson.D{}

	objectIds := make([]primitive.ObjectID, 0)
	for _, itemId := range req.Ids {
		objectIds = append(objectIds, utils.ConvertToObjectId(strings.TrimPrefix(itemId, "item:")))
	}

	filter = append(filter, bson.E{"_id", bson.D{{"$in", objectIds}}})

	results, err := u.itemRepository.FindManyItem(pctx, filter, nil)
	if err != nil {
		return nil, err
	}

	resultToRes := make([]*itemPb.Item, 0)
	for _, result := range results {
		resultToRes = append(resultToRes, &itemPb.Item{
			Id:       result.ItemId,
			Title:    result.Title,
			Price:    result.Price,
			ImageUrl: result.ImageUrl,
			Damage:   int32(result.Damage),
		})
	}

	return &itemPb.FindItemsInIdsRes{
		Items: resultToRes,
	}, nil

}
