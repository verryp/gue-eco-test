package repository

import (
	"github.com/verryp/gue-eco-test/internal/auth/common"
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
		//
	}
)
