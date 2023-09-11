package service

import (
	"context"
	"fmt"

	"github.com/verryp/gue-eco-test/internal/order/presentation"
	"github.com/verryp/gue-eco-test/internal/order/transformer"
)

type cartSvc struct {
	*Option
}

func NewCartService(opt *Option) ICartService {
	return &cartSvc{
		opt,
	}
}

func (svc *cartSvc) Add(ctx context.Context, req presentation.AddCartRequest) (err error) {
	respGetItem, err := svc.ProductAPI.GetByID(ctx, req.ItemID)
	if err != nil {
		svc.Log.Err(err).Msg("error get item id from product api")
		return err
	}

	if respGetItem.Data.IsAvailable != "yes" {
		errMsg := fmt.Sprintf("item with id %s is not available", req.ItemID)
		svc.Log.Warn().Msg(errMsg)
		return fmt.Errorf(errMsg)
	}

	item := presentation.ItemResponse{
		ID:             respGetItem.Data.ID,
		Name:           respGetItem.Data.Name,
		QuotaPerDays:   respGetItem.Data.QuotaPerDays,
		QuotaRemaining: respGetItem.Data.QuotaRemaining,
		Quantity:       respGetItem.Data.Quantity,
		Category:       respGetItem.Data.Category,
		IsAvailable:    respGetItem.Data.IsAvailable,
		Price:          respGetItem.Data.Price,
		CreatedAt:      respGetItem.Data.CreatedAt,
		UpdatedAt:      respGetItem.Data.UpdatedAt,
	}

	err = svc.Repository.Order.Create(ctx, transformer.ToParamCreateOrder(req, item))
	if err != nil {
		return
	}

	return
}
