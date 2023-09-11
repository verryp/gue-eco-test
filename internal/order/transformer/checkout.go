package transformer

import (
	"fmt"
	"time"

	"github.com/spf13/cast"
	"github.com/verryp/gue-eco-test/internal/order/consts"
	"github.com/verryp/gue-eco-test/internal/order/model"
	"github.com/verryp/gue-eco-test/internal/order/presentation"
	"github.com/verryp/gue-eco-test/pkg/generator"
)

func ConstructUpdateOrderForCheckout(ord model.Order) *model.ParamUpdateOrder {
	newOrd := &model.Order{
		ID:            ord.ID,
		OrderSerial:   ord.OrderSerial,
		CustomerName:  ord.CustomerName,
		CustomerEmail: ord.CustomerEmail,
		Status:        consts.OrderStatusUnpaid,
		TotalAmount:   ord.TotalAmount,
		ExpiredAt:     ord.ExpiredAt,
		CreatedAt:     ord.CreatedAt,
		UpdatedAt:     time.Now(),
	}

	return &model.ParamUpdateOrder{
		Order: newOrd,
		OrderHistory: &model.OrderHistory{
			ID:        cast.ToUint64(generator.GenerateInt64()),
			OrderID:   newOrd.ID,
			Status:    newOrd.Status,
			Remark:    fmt.Sprintf("order with serial %s is updated, previously %s", ord.OrderSerial, ord.Status),
			CreatedAt: time.Now(),
		},
	}
}

func ToCheckoutResponse(ord model.ParamUpdateOrder, detail model.OrderDetail) *presentation.CheckoutResponse {
	resp := &presentation.CheckoutResponse{
		ID:            cast.ToString(ord.Order.ID),
		Serial:        ord.Order.OrderSerial,
		CustomerName:  ord.Order.CustomerName,
		CustomerEmail: ord.Order.CustomerEmail,
		Status:        ord.Order.Status,
		TotalAmount:   ord.Order.TotalAmount,
		ExpiredAt:     ord.Order.ExpiredAt.Format(consts.DefaultLayoutDateTimeFormat),
		CreatedAt:     ord.Order.CreatedAt.Format(consts.DefaultLayoutDateTimeFormat),
		UpdatedAt:     ord.Order.UpdatedAt.Format(consts.DefaultLayoutDateTimeFormat),
		Detail: presentation.OrderDetailResponse{
			ID:           cast.ToInt(detail.ID),
			ItemID:       cast.ToString(detail.ItemID),
			ItemName:     detail.ItemName,
			ItemPrice:    detail.ItemPrice,
			Quantity:     detail.Quantity,
			TotalAmount:  detail.TotalAmount,
			CustomerNote: detail.CustomerNote,
		},
	}

	return resp
}

func ConstructUpdateOrderForCanceled(req presentation.CancelOrderRequest, ord model.Order) *model.ParamUpdateOrder {
	newOrd := &model.Order{
		ID:            ord.ID,
		OrderSerial:   ord.OrderSerial,
		CustomerName:  ord.CustomerName,
		CustomerEmail: ord.CustomerEmail,
		Status:        consts.OrderStatusCanceled,
		TotalAmount:   ord.TotalAmount,
		ExpiredAt:     ord.ExpiredAt,
		CreatedAt:     ord.CreatedAt,
		UpdatedAt:     time.Now(),
	}

	return &model.ParamUpdateOrder{
		Order: newOrd,
		OrderHistory: &model.OrderHistory{
			ID:        cast.ToUint64(generator.GenerateInt64()),
			OrderID:   newOrd.ID,
			Status:    newOrd.Status,
			Remark:    req.Reason,
			CreatedAt: time.Now(),
		},
	}
}
