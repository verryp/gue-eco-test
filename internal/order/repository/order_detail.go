package repository

import (
	"context"
	"database/sql"

	"github.com/verryp/gue-eco-test/internal/order/model"
)

type orderDetailRepo struct {
	*Option
}

func NewOrderDetailRepo(opt *Option) OrderDetailRepo {
	return &orderDetailRepo{
		opt,
	}
}

func (repo *orderDetailRepo) FindByOrderID(ctx context.Context, orderID string) (result *model.OrderDetail, err error) {
	result = &model.OrderDetail{}
	q := `
		SELECT
			id,
			order_id,
			item_id,
			item_name,
			item_price,
			quantity,
			item_price,
			total_amount,
			IFNULL(customer_note, '') customer_note,
			created_at,
			updated_at
		FROM
			order_details
		WHERE
			deleted_at IS NULL
			AND order_id=?
	`
	err = repo.DB.SelectOne(&result, q, orderID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		repo.Log.Err(err).Msg("FindByOrderID db error")
		return
	}

	return
}
