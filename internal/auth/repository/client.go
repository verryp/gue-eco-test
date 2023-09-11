package repository

import (
	"context"
	"database/sql"

	"github.com/verryp/gue-eco-test/internal/auth/model"
)

type clientRepo struct {
	*Option
}

func NewClientRepo(opt *Option) ClientRepo {
	return &clientRepo{
		opt,
	}
}

func (repo *clientRepo) FindByAPIKey(ctx context.Context, apiKey string) (cl *model.Client, err error) {
	cl = &model.Client{}
	q := `
		SELECT
			id, name, api_key, algorithm, location, public_cert, private_cert, created_at, updated_at
		FROM
			clients
		WHERE
			deleted_at IS NULL
			AND api_key=?
	`
	err = repo.DB.SelectOne(cl, q, apiKey)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		repo.Log.Err(err).Msg("FindByID item db error")
		return
	}

	return
}

func (repo *clientRepo) FindByID(ctx context.Context, id string) (cl *model.Client, err error) {
	cl = &model.Client{}
	q := `
		SELECT
			id, name, api_key, algorithm, location, public_cert, private_cert, created_at, updated_at
		FROM
			clients
		WHERE
			deleted_at IS NULL
			AND id=?
	`
	err = repo.DB.SelectOne(cl, q, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		repo.Log.Err(err).Msg("FindByID item db error")
		return
	}

	return
}
