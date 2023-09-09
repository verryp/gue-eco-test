package service

import (
	"context"

	"github.com/verryp/gue-eco-test/internal/auth/authenticator"
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/presentation"
	"github.com/verryp/gue-eco-test/internal/auth/repository"
)

type (
	Option struct {
		*common.Option
		Repository    *repository.Repository
		Authenticator authenticator.Authenticator
	}

	Service struct {
		HealthCheck IHealthCheck
		Auth        IAuthService
	}
)

type (
	IAuthService interface {
		Register(ctx context.Context, req presentation.SignUpRequest) error
		ClientAuthorization(ctx context.Context, apiKey, path string) (*presentation.GenerateClientTokenResponse, error)
		ValidateToken(ctx context.Context, token string) (*presentation.ValidateTokenResponse, error)
	}
)
