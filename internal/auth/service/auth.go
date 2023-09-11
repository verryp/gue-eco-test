package service

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cast"
	"github.com/verryp/gue-eco-test/internal/auth/consts"
	"github.com/verryp/gue-eco-test/internal/auth/helper"
	"github.com/verryp/gue-eco-test/internal/auth/model"
	"github.com/verryp/gue-eco-test/internal/auth/presentation"
	"github.com/verryp/gue-eco-test/pkg/generator"
)

type authSvc struct {
	*Option
}

func NewAuthService(opt *Option) IAuthService {
	return &authSvc{
		Option: opt,
	}
}

func (svc *authSvc) Register(ctx context.Context, req presentation.SignUpRequest) error {
	user, err := svc.Repository.User.FindByEmail(ctx, req.Email)
	if err != nil {
		svc.Log.Err(err).Msg("failed find user by email")
		return err
	}

	if user != nil {
		msg := "user already registered"
		svc.Log.Err(err).Msg(msg)
		return fmt.Errorf(msg)
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		svc.Log.Err(err).Msg("hashing password is failed")
		return err
	}

	err = svc.Repository.User.Create(ctx, &model.User{
		ID:        cast.ToUint64(generator.GenerateInt64()),
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		svc.Log.Err(err).Msg("failed create user")
		return err
	}

	return nil
}

func (svc *authSvc) ClientAuthorization(ctx context.Context, apiKey, pathURL string) (*presentation.GenerateClientTokenResponse, error) {
	token, cl, err := svc.Authenticator.GenerateClientToken(ctx, presentation.GenerateClientTokenRequest{
		Issuer: pathURL,
		APIKey: apiKey,
	})
	if err != nil {
		svc.Log.Err(err).Msg("failed generate client token")
		return nil, err
	}

	return &presentation.GenerateClientTokenResponse{
		TokenType:     consts.TokenTypeBearer,
		Client:        cl.Name,
		GuestToken:    token,
		ExpiredSecond: svc.Config.TokenExpiry.Guest,
	}, nil
}

func (svc *authSvc) ValidateToken(ctx context.Context, token string) (*presentation.ValidateTokenResponse, error) {
	validatedToken, err := svc.Authenticator.Authenticate(ctx, helper.SplitHeaderBearerToken(token))
	if err != nil {
		svc.Log.Err(err).Msg("failed authenticate token")
		return nil, err
	}

	jwtBlacklisted, err := svc.Cache.Get(ctx, helper.CacheKeyTokenBlacklisted(validatedToken.JwtID)).Bytes()
	if err != nil && err.Error() != consts.CacheNil {
		svc.Log.Err(err).Msg("error get date from cache")
		return nil, err
	}

	if string(jwtBlacklisted) != "" {
		msg := "the token is blacklisted"
		svc.Log.Err(err).Msg(msg)

		return nil, fmt.Errorf(msg)
	}

	return &presentation.ValidateTokenResponse{
		Client: validatedToken.Client,
		User:   validatedToken.User,
	}, nil
}

func (svc *authSvc) Login(ctx context.Context, pathURL string, req presentation.SignInRequest) (*presentation.SignInResponse, error) {
	user, err := svc.Repository.User.FindByEmail(ctx, req.Email)
	if err != nil {
		svc.Log.Err(err).Msg("failed find user by email")
		return nil, err
	}

	if user == nil {
		msg := "user not found"
		svc.Log.Error().Msg(msg)
		return nil, fmt.Errorf(msg)
	}

	err = helper.PasswordVerify(req.Password, user.Password)
	if err != nil {
		svc.Log.Err(err).Msg("verify password is failed")
		return nil, err
	}

	userToken, err := svc.Authenticator.GenerateAccessRefreshToken(ctx, presentation.GenerateAccessRefreshTokenRequest{
		Issuer: pathURL,
		User: presentation.UserPayload{
			ID:    cast.ToInt64(user.ID),
			Name:  user.Name,
			Email: user.Email,
		},
		Client: presentation.ClientPayload{
			ID: cast.ToInt(req.ClientID),
		},
	})
	if err != nil {
		svc.Log.Err(err).Msg("failed generate token")
		return nil, err
	}

	return &presentation.SignInResponse{
		TokenType: userToken.TokenType,
		AccessToken: presentation.TokenPayload{
			Value:     userToken.AccessToken.Value,
			ExpiredAt: userToken.AccessToken.ExpiredAt,
		},
		RefreshToken: presentation.TokenPayload{
			Value:     userToken.RefreshToken.Value,
			ExpiredAt: userToken.RefreshToken.ExpiredAt,
		},
	}, nil
}

func (svc *authSvc) RefreshToken(ctx context.Context, req presentation.ReTokenRequest) (*presentation.SignInResponse, error) {
	userToken, err := svc.Authenticator.GenerateAccessRefreshTokenByReToken(ctx, presentation.GenerateAccessRefreshTokenRequest{
		RefreshToken: helper.SplitHeaderBearerToken(req.Token),
		Issuer:       req.PathURL,
		Client: presentation.ClientPayload{
			ID: cast.ToInt(req.ClientID),
		},
	})
	if err != nil {
		svc.Log.Err(err).Msg("failed re-generate token")
		return nil, err
	}

	return &presentation.SignInResponse{
		TokenType: userToken.TokenType,
		AccessToken: presentation.TokenPayload{
			Value:     userToken.AccessToken.Value,
			ExpiredAt: userToken.AccessToken.ExpiredAt,
		},
		RefreshToken: presentation.TokenPayload{
			Value:     userToken.RefreshToken.Value,
			ExpiredAt: userToken.RefreshToken.ExpiredAt,
		},
	}, nil
}

func (svc *authSvc) BlackListToken(ctx context.Context, token string) (err error) {
	tokenAuth, err := svc.Authenticator.Authenticate(ctx, token)
	if err != nil {
		svc.Log.Err(err).Msg("failed authenticate token")
		return
	}

	if tokenAuth.User == nil {
		msg := "token is invalid"
		svc.Log.Error().Msg(msg)
		return fmt.Errorf(msg)
	}

	timeExpire := time.Unix(tokenAuth.ExpiredAt, 0)

	jwtID := tokenAuth.JwtID
	cmdStatus := svc.Cache.Set(ctx, helper.CacheKeyTokenBlacklisted(jwtID), jwtID, timeExpire.Sub(time.Now()))
	if cmdStatus.Err() != nil {
		svc.Log.Err(cmdStatus.Err()).Msg("failed set date to cache")
		return cmdStatus.Err()
	}

	return
}
