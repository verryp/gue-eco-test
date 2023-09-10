package service

import (
	"context"

	"github.com/verryp/gue-eco-test/internal/order/common"
	"github.com/verryp/gue-eco-test/internal/order/connector/product"
	"github.com/verryp/gue-eco-test/internal/order/presentation"
	"github.com/verryp/gue-eco-test/internal/order/repository"
)

type (
	Option struct {
		*common.Option
		Repository *repository.Repository
		ProductAPI product.API
	}

	Service struct {
		HealthCheck IHealthCheck
		Order       IOrderService
		Cart        ICartService
	}
)

type (
	ICartService interface {
		Add(ctx context.Context, req presentation.AddCartRequest) (err error)
	}

	IOrderService interface {
		Create(ctx context.Context, req presentation.CheckoutRequest) (resp *presentation.CheckoutResponse, err error)
		Cancel(ctx context.Context, req presentation.CancelOrderRequest) (*presentation.CheckoutResponse, error)
		List(ctx context.Context) (*presentation.OrderListResponses, error)
	}
)
