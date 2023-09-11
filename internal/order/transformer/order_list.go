package transformer

import (
	"github.com/spf13/cast"
	"github.com/verryp/gue-eco-test/internal/order/consts"
	"github.com/verryp/gue-eco-test/internal/order/model"
	"github.com/verryp/gue-eco-test/internal/order/presentation"
)

func ToOrderListResponse(ord model.Order) presentation.OrderListResponse {
	return presentation.OrderListResponse{
		ID:            cast.ToString(ord.ID),
		Serial:        ord.OrderSerial,
		CustomerName:  ord.CustomerName,
		CustomerEmail: ord.CustomerEmail,
		Status:        ord.Status,
		TotalAmount:   ord.TotalAmount,
		ExpiredAt:     ord.ExpiredAt.Format(consts.DefaultLayoutDateTimeFormat),
		CreatedAt:     ord.CreatedAt.Format(consts.DefaultLayoutDateTimeFormat),
		UpdatedAt:     ord.UpdatedAt.Format(consts.DefaultLayoutDateTimeFormat),
	}
}

func ToOrderListResponses(ords []model.Order) *presentation.OrderListResponses {
	results := &presentation.OrderListResponses{}

	for _, ord := range ords {
		results.Items = append(results.Items, ToOrderListResponse(ord))
	}

	return results
}
