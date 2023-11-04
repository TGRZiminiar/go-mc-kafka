package itemusecase

import (
	"context"
	"errors"

	"github.com/TGRZiminiar/go-mc-kafka/modules/item"
	"github.com/TGRZiminiar/go-mc-kafka/modules/item/itemRepository"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/utils"
)

type (
	ItemUsecaseService interface {
		CreateItem(pctx context.Context, req *item.CreateItemReq) (any, error)
		FindOneItem(pctx context.Context, itemId string) (*item.ItemShowCase, error)
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
		ItemId:   result.Id.Hex(),
		Title:    result.Title,
		Price:    result.Price,
		Damage:   result.Damage,
		ImageUrl: result.ImageUrl,
	}, nil
}
