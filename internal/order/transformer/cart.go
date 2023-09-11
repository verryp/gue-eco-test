package transformer

import (
	"fmt"
	"time"

	"github.com/spf13/cast"
	"github.com/verryp/gue-eco-test/internal/order/consts"
	"github.com/verryp/gue-eco-test/internal/order/helper"
	"github.com/verryp/gue-eco-test/internal/order/model"
	"github.com/verryp/gue-eco-test/internal/order/presentation"
	"github.com/verryp/gue-eco-test/pkg/generator"
)

func ToParamCreateOrder(req presentation.AddCartRequest, item presentation.ItemResponse) *model.ParamCreateOrder {
	var (
		result = &model.ParamCreateOrder{}
		amount = float64(req.Quantity) * item.Price
	)

	ord := &model.Order{
		ID:            cast.ToUint64(generator.GenerateInt64()),
		OrderSerial:   fmt.Sprintf("%s/ORD/%s", helper.GenerateRandomNumberString(6), time.Now().Format(consts.LayoutDateStripFormat)),
		CustomerName:  req.CustomerName,
		CustomerEmail: req.CustomerEmail,
		Status:        consts.OrderStatusOnCart,
		TotalAmount:   amount,
		ExpiredAt:     time.Now().Add(time.Hour * 24), // for scheduling to check the expiry order
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	history := &model.OrderHistory{
		ID:        cast.ToUint64(generator.GenerateInt64()),
		OrderID:   ord.ID,
		Status:    ord.Status,
		Remark:    fmt.Sprintf("order initiated, already on cart"),
		CreatedAt: time.Now(),
	}
	result.Order = ord
	result.OrderHistory = history
	result.OrderDetail = &model.OrderDetail{
		OrderID:      ord.ID,
		ItemID:       cast.ToInt64(req.ItemID),
		ItemName:     item.Name,
		ItemPrice:    item.Price,
		Quantity:     req.Quantity,
		TotalAmount:  amount,
		CustomerNote: req.CustomerNote,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return result
}
