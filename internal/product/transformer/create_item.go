package transformer

import (
	"time"

	"github.com/spf13/cast"
	"github.com/verryp/gue-eco-test/internal/product/model"
	"github.com/verryp/gue-eco-test/internal/product/presentation"
	"github.com/verryp/gue-eco-test/pkg/generator"
)

func TransformToCreateItem(req presentation.CreateItemRequest) *model.ParamCreateItem {
	item := &model.Item{
		ID:           cast.ToUint64(generator.GenerateInt64()),
		Name:         req.Name,
		QuotaPerDays: req.QuotaPerDays,
		Quantity:     req.Quantity,
		Category:     req.Category,
		Price:        req.Price,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	quota := &model.Quota{
		ItemID:         cast.ToInt64(item.ID),
		DateLimiter:    time.Now(),
		QuotaRemaining: req.Quantity,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return &model.ParamCreateItem{
		Item:  item,
		Quota: quota,
	}
}
