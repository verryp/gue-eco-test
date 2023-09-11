package transformer

import (
	"github.com/spf13/cast"
	"github.com/verryp/gue-eco-test/internal/product/consts"
	"github.com/verryp/gue-eco-test/internal/product/model"
	"github.com/verryp/gue-eco-test/internal/product/presentation"
)

func ToItemListResponse(item *model.ItemQuota) *presentation.ItemListResponse {
	isAvailable := "yes"
	if item.Quantity <= 0 || item.QuotaRemaining <= 0 {
		isAvailable = "no"
	}

	return &presentation.ItemListResponse{
		ID:             cast.ToString(item.ID),
		Name:           item.Name,
		QuotaRemaining: item.QuotaRemaining,
		Quantity:       item.Quantity,
		IsAvailable:    isAvailable,
		Category:       item.Category,
		Price:          item.Price,
		CreatedAt:      item.CreatedAt.Format(consts.DefaultLayoutDateTimeFormat),
		UpdatedAt:      item.UpdatedAt.Format(consts.DefaultLayoutDateTimeFormat),
	}
}

func ToItemListResponses(item []model.ItemQuota) presentation.ItemListResponses {
	var results []presentation.ItemListResponse

	for _, i := range item {
		results = append(results, *ToItemListResponse(&i))
	}

	return presentation.ItemListResponses{
		Items: results,
	}
}
