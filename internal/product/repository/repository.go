package repository

import (
	"context"

	"github.com/verryp/gue-eco-test/internal/product/common"
	"github.com/verryp/gue-eco-test/internal/product/model"
	"gopkg.in/gorp.v2"
)

type (
	Option struct {
		*common.Option
		DB *gorp.DbMap
	}

	Repository struct {
		Item ItemRepo
	}
)

type (
	ItemRepo interface {
		Fetch(ctx context.Context) (items []model.Item, count int64, err error)
		FindByID(ctx context.Context, id string) (item *model.Item, err error)
		Create(ctx context.Context, item *model.Item) error
		Update(ctx context.Context, item *model.Item) (int64, error)
	}
)
