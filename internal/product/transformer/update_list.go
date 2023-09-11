package transformer

import (
	"fmt"
	"time"

	"github.com/verryp/gue-eco-test/internal/product/model"
	"github.com/verryp/gue-eco-test/internal/product/presentation"
)

func ConstructParamUpdateItem(r presentation.UpdateItemRequest, item model.ItemQuota, quota model.Quota) *model.ParamCreateItem {
	if r.Name != "" {
		item.Name = r.Name
	}

	if r.Category != "" {
		item.Category = r.Category
	}

	if r.Quantity != 0 {
		item.Quantity = r.Quantity
	}

	if r.QuotaPerDays != 0 {
		item.QuotaPerDays = r.QuotaPerDays

		// note: actively renew the `date_limiter` if daily quotas has changes
		tn := time.Now()
		quota.DateLimiter = tn
		quota.UpdatedAt = tn
	}

	if r.Price != 0 {
		item.Price = r.Price
	}

	newItem := &model.Item{
		ID:           item.ID,
		Name:         item.Name,
		QuotaPerDays: item.QuotaPerDays,
		Quantity:     item.Quantity,
		Category:     item.Category,
		Price:        item.Price,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    time.Now(),
	}

	return &model.ParamCreateItem{
		Item:  newItem,
		Quota: &quota,
	}
}

func ConstructDecreaseItemQuantity(r presentation.UpdateItemRequest, item model.ItemQuota, quota model.Quota) (*model.ParamCreateItem, error) {
	quota.QuotaRemaining = quota.QuotaRemaining - r.Quantity
	if quota.QuotaRemaining < 0 {
		return nil, fmt.Errorf("the quota remaining is invalid")
	}

	quota.UpdatedAt = time.Now()

	qty := item.Quantity - r.Quantity
	if qty < 0 {
		return nil, fmt.Errorf("the quantity remaining is invalid")
	}

	return &model.ParamCreateItem{
		Item: &model.Item{
			ID:           item.ID,
			Name:         item.Name,
			QuotaPerDays: item.QuotaPerDays,
			Quantity:     qty,
			Category:     item.Category,
			Price:        item.Price,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    time.Now(),
		},
		Quota: &quota,
	}, nil
}

func ConstructIncreaseItemQuantity(r presentation.UpdateItemRequest, item model.ItemQuota, quota model.Quota) (*model.ParamCreateItem, error) {
	quota.QuotaRemaining = quota.QuotaRemaining + r.Quantity
	quota.UpdatedAt = time.Now()

	return &model.ParamCreateItem{
		Item: &model.Item{
			ID:           item.ID,
			Name:         item.Name,
			QuotaPerDays: item.QuotaPerDays,
			Quantity:     item.Quantity + r.Quantity,
			Category:     item.Category,
			Price:        item.Price,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    time.Now(),
		},
		Quota: &quota,
	}, nil
}
