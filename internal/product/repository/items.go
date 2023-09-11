package repository

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/verryp/gue-eco-test/internal/product/model"
)

type itemsRepo struct {
	*Option
}

type RequestPagination struct {
	Query  string
	Sort   []string
	Limit  int
	Offset int
}

func NewItemsRepo(opt *Option) ItemRepo {
	return &itemsRepo{
		opt,
	}
}

func (repo *itemsRepo) Fetch(ctx context.Context) (items []model.ItemQuota, count int64, err error) {
	var (
		psql         = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
		queryBuilder = psql.
				Select().
				From("items i").
				Join("item_quotas iq").
				Where("i.deleted_at IS NULL AND iq.deleted_at IS NULL")

		queryPagination = queryBuilder.
				Columns("i.id id, i.name, iq.date_limiter, iq.quota_remaining, i.quota_per_days, i.quantity, i.category, i.price, i.created_at, i.updated_at")
	)

	query, args, err := queryPagination.ToSql()
	if err != nil {
		repo.Log.Err(err).Msgf("Fetch items pagination toSql error")
		return
	}

	_, err = repo.DB.Select(&items, query, args...)
	if err != nil {
		repo.Log.Err(err).Msgf("Fetch items pagination select error")
	}

	queryCount := queryBuilder.Columns("COUNT(i.id)")

	countSql, args, err := queryCount.ToSql()
	if err != nil {
		repo.Log.Err(err).Msgf("Fetch items count toSql error")
		return
	}

	count, err = repo.DB.SelectInt(countSql, args...)
	if err != nil {
		repo.Log.Err(err).Msgf("Fetch items count select error")
		return
	}

	return items, count, err
}

func (repo *itemsRepo) FindByID(ctx context.Context, id string) (item *model.ItemQuota, err error) {
	item = &model.ItemQuota{}
	q := `
		SELECT
			i.id, i.name, iq.date_limiter, iq.quota_remaining, i.quota_per_days, i.quantity, i.category, i.price, i.created_at, i.updated_at
		FROM
			items i
		JOIN
			item_quotas iq
		WHERE
			i.deleted_at IS NULL
			AND iq.deleted_at IS NULL
			AND i.id=?
	`
	err = repo.DB.SelectOne(item, q, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		repo.Log.Err(err).Msg("FindByID item db error")
		return
	}

	return
}

func (repo *itemsRepo) Create(ctx context.Context, p *model.ParamCreateItem) (err error) {
	tx, err := repo.DB.Begin()
	if err != nil {
		repo.Log.Err(err).Msg("failed begin tx")
		return
	}

	if err = tx.Insert(p.Item); err != nil {
		repo.Log.Err(err).Msg("failed create item")
		return tx.Rollback()
	}

	if err = tx.Insert(p.Quota); err != nil {
		repo.Log.Err(err).Msg("failed create item quota")
		return tx.Rollback()
	}

	if err = tx.Commit(); err != nil {
		repo.Log.Err(err).Msg("failed commit create item")
		return
	}

	return nil
}

func (repo *itemsRepo) Update(ctx context.Context, item *model.ParamCreateItem) (int64, error) {
	tx, err := repo.DB.Begin()
	if err != nil {
		repo.Log.Err(err).Msg("begin repo is failed")
		return 0, nil
	}

	rows, err := repo.DB.Update(item.Item)
	if err != nil {
		repo.Log.Err(err).Msg("error update item")
		return 0, tx.Rollback()
	}

	if rows <= 0 {
		repo.Log.Warn().Msg("no rows is updated")
		return 0, tx.Rollback()
	}

	if item.Quota != nil {
		_, err := repo.DB.Update(item.Quota)
		if err != nil {
			return 0, tx.Rollback()
		}
	}

	err = tx.Commit()
	if err != nil {
		repo.Log.Err(err).Msg("failed commit update item")
		return 0, err
	}

	return rows, nil
}
