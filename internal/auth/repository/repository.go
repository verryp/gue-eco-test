package repository

import (
	"context"

	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/model"
	"gopkg.in/gorp.v2"
)

type (
	Option struct {
		*common.Option
		DB *gorp.DbMap
	}

	Repository struct {
		User UserRepo
	}
)

type (
	UserRepo interface {
		Create(ctx context.Context, user *model.User) error
	}
)
