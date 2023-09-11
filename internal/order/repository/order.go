package repository

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/verryp/gue-eco-test/internal/order/model"
)

type orderRepo struct {
	*Option
}

func NewOrderRepo(opt *Option) OrderRepo {
	return &orderRepo{
		opt,
	}
}

func (repo *orderRepo) Create(ctx context.Context, p *model.ParamCreateOrder) (err error) {
	tx, err := repo.DB.Begin()
	if err != nil {
		repo.Log.Err(err).Msg("failed begin tx")
		return
	}

	err = tx.Insert(p.Order)
	if err != nil {
		repo.Log.Err(err).Msg("failed create order")
		return tx.Rollback()
	}

	err = tx.Insert(p.OrderHistory)
	if err != nil {
		repo.Log.Err(err).Msg("failed create order history")
		return tx.Rollback()
	}

	err = tx.Insert(p.OrderDetail)
	if err != nil {
		repo.Log.Err(err).Msg("failed create order detail")
		return tx.Rollback()
	}

	err = tx.Commit()
	if err != nil {
		repo.Log.Err(err).Msg("failed commit order")
		return
	}

	return
}

func (repo *orderRepo) Update(ctx context.Context, p *model.ParamUpdateOrder) (err error) {
	tx, err := repo.DB.Begin()
	if err != nil {
		repo.Log.Err(err).Msg("failed begin tx")
		return
	}

	_, err = tx.Update(p.Order)
	if err != nil {
		repo.Log.Err(err).Msg("failed update order")
		return tx.Rollback()
	}

	err = tx.Insert(p.OrderHistory)
	if err != nil {
		repo.Log.Err(err).Msg("failed create order history")
		return tx.Rollback()
	}

	err = tx.Commit()
	if err != nil {
		repo.Log.Err(err).Msg("failed commit order")
		return
	}

	return
}

func (repo *orderRepo) FindBy(ctx context.Context, id, status string) (item *model.Order, err error) {
	item = &model.Order{}
	q := `
		SELECT
			id,
			order_serial,
			IFNULL(customer_name, '') customer_name,
			IFNULL(customer_email, '') customer_email,
			status,
			total_amount,
			expired_at,
			created_at,
			updated_at
		FROM
			orders
		WHERE
			deleted_at IS NULL
			AND id=?
			AND status=?
		LIMIT 1
	`

	err = repo.DB.SelectOne(item, q, id, status)
	if err == sql.ErrNoRows {
		repo.Log.Err(err).Msg("FindByID item is empty")
		return nil, nil
	}

	if err != nil {
		repo.Log.Err(err).Msg("FindByID item db error")
		return
	}

	return
}

func (repo *orderRepo) Fetch(ctx context.Context) (items []model.Order, count int64, err error) {
	var (
		psql            = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
		queryBuilder    = psql.Select().From("orders").Where("deleted_at IS NULL")
		queryPagination = queryBuilder.Columns(
			`id,
			order_serial,
			IFNULL(customer_name, '') customer_name,
			IFNULL(customer_email, '') customer_email,
			status,
			total_amount,
			expired_at,
			created_at,
			updated_at`,
		)
	)

	query, args, err := queryPagination.ToSql()
	if err != nil {
		repo.Log.Err(err).Msgf("Fetch orders pagination toSql error")
		return
	}

	_, err = repo.DB.Select(&items, query, args...)
	if err != nil {
		repo.Log.Err(err).Msgf("Fetch order pagination select error")
	}

	queryCount := queryBuilder.Columns("COUNT(id)")

	countSql, args, err := queryCount.ToSql()
	if err != nil {
		repo.Log.Err(err).Msgf("Fetch order count toSql error")
		return
	}

	count, err = repo.DB.SelectInt(countSql, args...)
	if err != nil {
		repo.Log.Err(err).Msgf("Fetch order count select error")
		return
	}

	return items, count, err
}
