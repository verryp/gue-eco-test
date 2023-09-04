package service

import (
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/repository"
)

type (
	Option struct {
		*common.Option
		Repository *repository.Repository
	}

	Service struct {
		HealthCheck IHealthCheck
		User        IUserService
	}
)

type (
	IUserService interface {
		//
	}
)
