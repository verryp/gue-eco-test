package repository

import (
	"context"
	"database/sql"

	"github.com/verryp/gue-eco-test/internal/auth/model"
)

type userRepo struct {
	*Option
}

func NewUserRepo(opt *Option) UserRepo {
	return &userRepo{
		opt,
	}
}

func (repo *userRepo) Create(ctx context.Context, user *model.User) error {
	return repo.DB.Insert(user)
}

func (repo *userRepo) FindByID(ctx context.Context, id string) (cl *model.User, err error) {
	cl = &model.User{}
	q := `
		SELECT
			id, name, email, created_at, updated_at
		FROM
			users
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

func (repo *userRepo) FindByEmail(ctx context.Context, email string) (cl *model.User, err error) {
	cl = &model.User{}
	q := `
		SELECT
			id, name, email, password, created_at, updated_at
		FROM
			users
		WHERE
			deleted_at IS NULL
			AND email=?
	`
	err = repo.DB.SelectOne(cl, q, email)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		repo.Log.Err(err).Msg("FindByID item db error")
		return
	}

	return
}
