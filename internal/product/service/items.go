package service

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cast"
	"github.com/verryp/gue-eco-test/internal/product/presentation"
	"github.com/verryp/gue-eco-test/internal/product/transformer"
)

type itemSvc struct {
	*Option
}

func NewItemService(opt *Option) IItemService {
	return &itemSvc{
		Option: opt,
	}
}

func (svc *itemSvc) List(ctx context.Context) (resp presentation.ItemListResponses, err error) {
	items, _, err := svc.Repository.Item.Fetch(ctx)
	if err != nil {
		return
	}

	resp = transformer.ToItemListResponses(items)
	return
}

func (svc *itemSvc) FindByID(ctx context.Context, id string) (resp *presentation.ItemListResponse, err error) {
	item, err := svc.Repository.Item.FindByID(ctx, id)
	if err != nil {
		svc.Log.Err(err).Msg("error find item db")
		return
	}

	if item == nil {
		svc.Log.Warn().Msg("item is not found")
		return
	}

	resp = transformer.ToItemListResponse(item)
	return
}

func (svc *itemSvc) Create(ctx context.Context, req presentation.CreateItemRequest) error {
	return svc.Repository.Item.Create(ctx, transformer.TransformToCreateItem(req))
}

func (svc *itemSvc) UpdateByID(ctx context.Context, id string, req presentation.UpdateItemRequest) (rows int64, err error) {
	item, err := svc.Repository.Item.FindByID(ctx, id)
	if err != nil {
		svc.Log.Err(err).Msg("error find item db")
		return 0, err
	}

	if item == nil {
		svc.Log.Warn().Msg("item not found")
		return 0, nil
	}

	itemQuota, err := svc.Repository.ItemQuota.FindByItemID(ctx, cast.ToString(item.ID))
	if err != nil {
		svc.Log.Err(err).Msg("failed to get item quota repo")
		return 0, err
	}

	if itemQuota == nil {
		msg := "item quota not found"
		svc.Log.Warn().Msg(msg)
		return 0, fmt.Errorf(msg)
	}

	return svc.Repository.Item.Update(ctx, transformer.ConstructParamUpdateItem(req, *item, *itemQuota))
}

func (svc *itemSvc) DecreaseItemQuantity(ctx context.Context, id string, req presentation.UpdateItemRequest) (rows int64, err error) {
	item, err := svc.Repository.Item.FindByID(ctx, id)
	if err != nil {
		svc.Log.Err(err).Msg("error find item db")
		return 0, err
	}

	if item == nil {
		svc.Log.Warn().Msg("item not found")
		return 0, nil
	}

	itemQuota, err := svc.Repository.ItemQuota.FindByItemID(ctx, cast.ToString(item.ID))
	if err != nil {
		svc.Log.Err(err).Msg("failed to get item quota repo")
		return
	}

	if itemQuota == nil {
		msg := "item quota not found"
		svc.Log.Warn().Msg(msg)

		return 0, fmt.Errorf(msg)
	}

	// renewing the quota
	if item.DateLimiter.Before(time.Now()) {
		tn := time.Now()
		itemQuota.DateLimiter = tn
		itemQuota.UpdatedAt = tn
		itemQuota.QuotaRemaining = item.QuotaPerDays
		_, err = svc.Repository.ItemQuota.Update(ctx, itemQuota)
		if err != nil {
			svc.Log.Err(err).Msg("failed to update item quota repo")
			return
		}
	}

	p, err := transformer.ConstructDecreaseItemQuantity(req, *item, *itemQuota)
	if err != nil {
		return
	}

	return svc.Repository.Item.Update(ctx, p)
}

func (svc *itemSvc) IncreaseItemQuantity(ctx context.Context, id string, req presentation.UpdateItemRequest) (rows int64, err error) {
	item, err := svc.Repository.Item.FindByID(ctx, id)
	if err != nil {
		svc.Log.Err(err).Msg("error find item db")
		return 0, err
	}

	itemQuota, err := svc.Repository.ItemQuota.FindByItemID(ctx, cast.ToString(item.ID))
	if err != nil {
		svc.Log.Err(err).Msg("failed to get item quota repo")
		return
	}

	if itemQuota == nil {
		msg := "item quota not found"
		svc.Log.Warn().Msg(msg)

		return 0, fmt.Errorf(msg)
	}

	p, err := transformer.ConstructIncreaseItemQuantity(req, *item, *itemQuota)
	if err != nil {
		return
	}

	return svc.Repository.Item.Update(ctx, p)
}
