package repository

import (
	"context"

	"github.com/verryp/gue-eco-test/internal/order/common"
	"github.com/verryp/gue-eco-test/internal/order/model"
	"gopkg.in/gorp.v2"
)

type (
	Option struct {
		*common.Option
		DB *gorp.DbMap
	}

	Repository struct {
		Order       OrderRepo
		OrderDetail OrderDetailRepo
	}
)

type (
	OrderRepo interface {
		Create(ctx context.Context, p *model.ParamCreateOrder) (err error)
		Update(ctx context.Context, p *model.ParamUpdateOrder) (err error)
		FindBy(ctx context.Context, id, status string) (item *model.Order, err error)
		Fetch(ctx context.Context) (items []model.Order, count int64, err error)
	}

	OrderDetailRepo interface {
		FetchByOrderID(ctx context.Context, orderID string) (results []model.OrderDetail, err error)
	}
)
