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
		Item      ItemRepo
		ItemQuota ItemQuotaRepo
	}
)

type (
	ItemRepo interface {
		Fetch(ctx context.Context) (items []model.ItemQuota, count int64, err error)
		FindByID(ctx context.Context, id string) (item *model.ItemQuota, err error)
		Create(ctx context.Context, p *model.ParamCreateItem) (err error)
		Update(ctx context.Context, item *model.ParamCreateItem) (int64, error)
	}

	ItemQuotaRepo interface {
		FindByItemID(ctx context.Context, itemID string) (quota *model.Quota, err error)
		Update(ctx context.Context, quota *model.Quota) (int64, error)
	}
)
