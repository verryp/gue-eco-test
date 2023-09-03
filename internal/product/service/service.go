package service

import (
	"context"

	"github.com/verryp/gue-eco-test/internal/product/common"
	"github.com/verryp/gue-eco-test/internal/product/presentation"
	"github.com/verryp/gue-eco-test/internal/product/repository"
)

type (
	Option struct {
		*common.Option
		Repository *repository.Repository
	}

	Service struct {
		HealthCheck IHealthCheck
		Item        IItemService
	}
)

type (
	IItemService interface {
		List(ctx context.Context) (items presentation.ItemListResponses, err error)
		FindByID(ctx context.Context, id string) (resp *presentation.ItemListResponse, err error)
		Create(ctx context.Context, req presentation.CreateItemRequest) error
		UpdateByID(ctx context.Context, id string, req presentation.UpdateItemRequest) (rows int64, err error)
	}
)
