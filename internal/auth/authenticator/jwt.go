package authenticator

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hashicorp/go-uuid"
	"github.com/spf13/cast"
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/consts"
	"github.com/verryp/gue-eco-test/internal/auth/helper"
	"github.com/verryp/gue-eco-test/internal/auth/presentation"
	"github.com/verryp/gue-eco-test/internal/auth/repository"
)

type jwtAuthenticator struct {
	*common.Option
	repo *repository.Repository
}

func NewJWTAuthenticator(opt *common.Option, repo *repository.Repository) Authenticator {
	return &jwtAuthenticator{
		Option: opt,
		repo:   repo,
	}
}

func (j *jwtAuthenticator) GenerateClientToken(ctx context.Context, req presentation.GenerateClientTokenRequest) (signedToken string, client *presentation.ClientPayload, err error) {
	cl, err := j.repo.Client.FindByAPIKey(ctx, req.APIKey)
	if err != nil {
		return
	}

	if cl == nil {
		return "", nil, fmt.Errorf("client not found")
	}

	keyData, err := os.ReadFile(fmt.Sprintf("%s/%s", cl.Location, cl.PrivateCert))
	if err != nil {
		return
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		j.Log.Err(err).Msg("err parse RSA private key")
		return
	}

	jti, _ := uuid.GenerateUUID()
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(j.Config.TokenExpiry.Guest)))

	method := jwt.GetSigningMethod(cl.Algorithm)
	token := jwt.NewWithClaims(method, presentation.StandardClaimRequest{
		Client:    cl.Name,
		TokenType: consts.TokenTypeGuestToken,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: req.Issuer,
			Audience: jwt.ClaimStrings{
				cast.ToString(cl.ID),
			},
			ExpiresAt: expiresAt,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        jti,
		},
	})

	t, err := token.SignedString(signKey)
	if err != nil {
		return
	}

	client = &presentation.ClientPayload{
		ID:   cl.ID,
		Name: cl.Name,
	}

	return t, client, nil
}

func (j *jwtAuthenticator) GenerateAccessRefreshToken(ctx context.Context, req presentation.GenerateAccessRefreshTokenRequest) (resp *presentation.GenerateAccessRefreshTokenResponse, err error) {
	cl, err := j.repo.Client.FindByID(ctx, cast.ToString(req.Client.ID))
	if err != nil {
		return
	}

	if cl == nil {
		return nil, fmt.Errorf("client not found")
	}

	keyData, err := os.ReadFile(fmt.Sprintf("%s/%s", cl.Location, cl.PrivateCert))
	if err != nil {
		return
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		j.Log.Err(err).Msg("err parse RSA private key")
		return
	}

	jti, _ := uuid.GenerateUUID()
	accessTTL := jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(j.Config.TokenExpiry.Access)))
	refreshTTL := jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(j.Config.TokenExpiry.Refresh)))

	method := jwt.GetSigningMethod(cl.Algorithm)

	accessClaim := presentation.StandardClaimRequest{
		Name:      req.User.Name,
		Email:     req.User.Email,
		Client:    cl.Name,
		TokenType: consts.TokenTypeAccessToken,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: req.Issuer,
			Audience: jwt.ClaimStrings{
				cast.ToString(cl.ID),
			},
			ExpiresAt: accessTTL,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        jti,
			Subject:   cast.ToString(req.User.ID),
		},
	}
	accessToken, err := jwt.NewWithClaims(method, accessClaim).SignedString(signKey)
	if err != nil {
		return
	}

	refreshClaim := accessClaim
	refreshClaim.ExpiresAt = refreshTTL
	refreshClaim.TokenType = consts.TokenTypeRefreshToken

	refreshToken, err := jwt.NewWithClaims(method, refreshClaim).SignedString(signKey)
	if err != nil {
		return
	}

	return &presentation.GenerateAccessRefreshTokenResponse{
		JwtID:     jti,
		UserName:  accessClaim.Name,
		UserEmail: accessClaim.Email,
		TokenType: consts.TokenTypeBearer,
		Audience:  cast.ToString(cl.ID),
		Subject:   accessClaim.Subject,
		Issuer:    req.Issuer,
		Client:    req.Client,
		RefreshToken: presentation.TokenPayload{
			Value:     refreshToken,
			ExpiredAt: refreshTTL.Unix(),
		},
		AccessToken: presentation.TokenPayload{
			Value:     accessToken,
			ExpiredAt: accessTTL.Unix(),
		},
	}, nil
}

func (j *jwtAuthenticator) GenerateAccessRefreshTokenByReToken(ctx context.Context, req presentation.GenerateAccessRefreshTokenRequest) (resp *presentation.GenerateAccessRefreshTokenResponse, err error) {
	cl, err := j.repo.Client.FindByID(ctx, cast.ToString(req.Client.ID))
	if err != nil {
		return
	}

	if cl == nil {
		return nil, fmt.Errorf("client not found")
	}

	retoken, err := j.ParseUnverified(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	if retoken.TokenType != consts.TokenTypeRefreshToken {
		return nil, fmt.Errorf("token type is not valid")
	}

	keyData, err := os.ReadFile(fmt.Sprintf("%s/%s", cl.Location, cl.PrivateCert))
	if err != nil {
		return
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		j.Log.Err(err).Msg("err parse RSA private key")
		return
	}

	var (
		jti       = retoken.ID
		accessTTL = jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(j.Config.TokenExpiry.Access)))
		method    = jwt.GetSigningMethod(cl.Algorithm)
		aud, _    = retoken.GetAudience()
	)

	accessClaim := presentation.StandardClaimRequest{
		Name:      retoken.Name,
		Email:     retoken.Email,
		Client:    retoken.Client,
		TokenType: consts.TokenTypeAccessToken,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    req.Issuer,
			Audience:  aud,
			ExpiresAt: accessTTL,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        jti,
			Subject:   retoken.Subject,
		},
	}
	accessToken, err := jwt.NewWithClaims(method, accessClaim).SignedString(signKey)
	if err != nil {
		return
	}

	return &presentation.GenerateAccessRefreshTokenResponse{
		JwtID:     jti,
		UserName:  accessClaim.Name,
		UserEmail: accessClaim.Email,
		TokenType: consts.TokenTypeBearer,
		Audience:  cast.ToString(cl.ID),
		Subject:   accessClaim.Subject,
		Issuer:    req.Issuer,
		Client:    req.Client,
		RefreshToken: presentation.TokenPayload{
			Value:     req.RefreshToken,
			ExpiredAt: retoken.ExpiresAt.Unix(),
		},
		AccessToken: presentation.TokenPayload{
			Value:     accessToken,
			ExpiredAt: accessTTL.Unix(),
		},
	}, nil
}

func (j *jwtAuthenticator) ParseUnverified(ctx context.Context, token string) (*presentation.StandardClaimRequest, error) {
	tokenParsed, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed claims parsed token")
	}

	var (
		iss, _ = claims.GetIssuer()
		aud, _ = claims.GetAudience()
		exp, _ = claims.GetExpirationTime()
		iat, _ = claims.GetIssuedAt()
		sub, _ = claims.GetSubject()
	)

	payload := &presentation.StandardClaimRequest{
		Client:    claims["client"].(string),
		TokenType: claims["token_type"].(string),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    iss,
			Subject:   sub,
			Audience:  aud,
			ExpiresAt: exp,
			IssuedAt:  iat,
			ID:        claims["jti"].(string),
		},
	}

	if name, ok := claims["name"]; ok {
		payload.Name = name.(string)
	}

	if email, ok := claims["email"]; ok {
		payload.Email = email.(string)
	}

	return payload, nil

}

func (j *jwtAuthenticator) Authenticate(ctx context.Context, token string) (resp *presentation.TokenAuthenticatedResponse, err error) {
	jwtClaim, err := j.ParseUnverified(ctx, token)
	if err != nil {
		return
	}

	var aud []string
	b, _ := jwtClaim.Audience.MarshalJSON()
	_ = json.Unmarshal(b, &aud)

	cl, err := j.repo.Client.FindByID(ctx, aud[0])
	if err != nil {
		return
	}

	if cl == nil {
		return nil, fmt.Errorf("client is not found")
	}

	pubCert, err := os.ReadFile(fmt.Sprintf("%s/%s", cl.Location, cl.PublicCert))
	if err != nil {
		return
	}
	isValidToken, err := j.validatePubKey(token, pubCert)
	if !isValidToken {
		return nil, err
	}

	var authUser *presentation.AuthenticatedUser
	if helper.InArray(jwtClaim.TokenType, []string{consts.TokenTypeAccessToken, consts.TokenTypeRefreshToken}) {
		authUser, err = j.getAuthenticatedUser(ctx, *jwtClaim)
		if err != nil {
			return nil, err
		}
	}

	return &presentation.TokenAuthenticatedResponse{
		JwtID:     jwtClaim.ID,
		IssuedAt:  jwtClaim.IssuedAt.Time,
		TokenType: jwtClaim.TokenType,
		ExpiredAt: jwtClaim.ExpiresAt.Unix(),
		Client: presentation.ClientPayload{
			ID:   cl.ID,
			Name: cl.Name,
		},
		User: authUser,
	}, nil
}

func (j *jwtAuthenticator) validatePubKey(token string, pubKey []byte) (bool, error) {
	verifiedKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		j.Log.Err(err).Msg("failed parse public key")
		return false, err
	}

	tokenParsed, err := jwt.ParseWithClaims(token, &presentation.StandardClaimRequest{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return verifiedKey, nil
	})
	if err != nil {
		return false, err
	}

	if !tokenParsed.Valid {
		return false, fmt.Errorf("token is invalid")
	}

	return true, nil

}

func (j *jwtAuthenticator) getAuthenticatedUser(ctx context.Context, r presentation.StandardClaimRequest) (*presentation.AuthenticatedUser, error) {
	u, err := j.repo.User.FindByID(ctx, r.Subject)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, fmt.Errorf("user not found")
	}

	return &presentation.AuthenticatedUser{
		ID:        cast.ToString(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Format(consts.DefaultLayoutDateTimeFormat),
		UpdatedAt: u.UpdatedAt.Format(consts.DefaultLayoutDateTimeFormat),
	}, nil
}
