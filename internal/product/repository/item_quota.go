package repository

import (
	"context"
	"database/sql"

	"github.com/verryp/gue-eco-test/internal/product/model"
)

type itemQuotaRepo struct {
	*Option
}

func NewItemQuotaRepo(opt *Option) ItemQuotaRepo {
	return &itemQuotaRepo{
		opt,
	}
}

func (repo *itemQuotaRepo) FindByItemID(ctx context.Context, itemID string) (quota *model.Quota, err error) {
	quota = &model.Quota{}
	q := `
		SELECT
			id, item_id, date_limiter, quota_remaining, created_at, updated_at
		FROM
			item_quotas
		WHERE
			deleted_at IS NULL
			AND item_id=?
	`
	err = repo.DB.SelectOne(quota, q, itemID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		repo.Log.Err(err).Msg("FindByItemID item quota db error")
		return
	}

	return
}

func (repo *itemQuotaRepo) Update(ctx context.Context, quota *model.Quota) (int64, error) {
	return repo.DB.Update(quota)
}
