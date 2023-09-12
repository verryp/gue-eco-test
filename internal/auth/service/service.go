package service

import (
	"context"

	"github.com/redis/go-redis/v9"
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
		Cache         *redis.Client
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
		ValidateToken(ctx context.Context, req presentation.ValidateTokenRequest) (*presentation.ValidateTokenResponse, error)
		Login(ctx context.Context, pathURL string, req presentation.SignInRequest) (*presentation.SignInResponse, error)
		RefreshToken(ctx context.Context, req presentation.ReTokenRequest) (*presentation.SignInResponse, error)
		BlackListToken(ctx context.Context, token string) (err error)
	}
)
