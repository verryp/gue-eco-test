package authenticator

import (
	"context"

	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/presentation"
)

type (
	Option struct {
		*common.Config
	}

	Authenticator interface {
		GenerateClientToken(ctx context.Context, req presentation.GenerateClientTokenRequest) (signedToken string, client *presentation.ClientPayload, err error)
		GenerateAccessRefreshToken(ctx context.Context, req presentation.GenerateAccessRefreshTokenRequest) (resp *presentation.GenerateAccessRefreshTokenResponse, err error)
		GenerateAccessRefreshTokenByReToken(ctx context.Context, req presentation.GenerateAccessRefreshTokenRequest) (resp *presentation.GenerateAccessRefreshTokenResponse, err error)
		ParseUnverified(ctx context.Context, token string) (*presentation.StandardClaimRequest, error)
		Authenticate(ctx context.Context, token string) (resp *presentation.TokenAuthenticatedResponse, err error)
	}
)
