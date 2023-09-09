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
		User   UserRepo
		Client ClientRepo
	}
)

type (
	UserRepo interface {
		Create(ctx context.Context, user *model.User) error
		FindByID(ctx context.Context, id string) (cl *model.User, err error)
		FindByEmail(ctx context.Context, email string) (cl *model.User, err error)
	}

	ClientRepo interface {
		FindByAPIKey(ctx context.Context, apiKey string) (cl *model.Client, err error)
		FindByID(ctx context.Context, id string) (cl *model.Client, err error)
	}
)
