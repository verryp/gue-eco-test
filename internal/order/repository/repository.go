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
		Order OrderRepo
	}
)

type (
	OrderRepo interface {
		Create(ctx context.Context, p *model.ParamCreateOrder) (err error)
	}
)
