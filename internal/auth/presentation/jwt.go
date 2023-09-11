package presentation

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	GenerateClientTokenRequest struct {
		Issuer string `json:"issuer"`
		APIKey string `json:"api_key"`
	}

	GenerateClientTokenResponse struct {
		TokenType     string `json:"token_type"`
		Client        string `json:"client"`
		ExpiredSecond int64  `json:"expired_second"`
		GuestToken    string `json:"guest_token"`
	}

	StandardClaimRequest struct {
		Name      string `json:"name,omitempty"`
		Email     string `json:"email,omitempty"`
		Client    string `json:"client,omitempty"`
		TokenType string `json:"token_type,omitempty"`
		jwt.RegisteredClaims
	}

	AuthenticatedUser struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	TokenAuthenticatedResponse struct {
		JwtID     string             `json:"jwt_id"`
		IssuedAt  time.Time          `json:"issued_at"`
		TokenType string             `json:"token_type"`
		ExpiredAt int64              `json:"expired_at,omitempty"`
		Client    ClientPayload      `json:"client"`
		User      *AuthenticatedUser `json:"user,omitempty"`
	}

	ClientPayload struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	UserPayload struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	GenerateAccessRefreshTokenRequest struct {
		RefreshToken string        `json:"refresh_token,omitempty"`
		Issuer       string        `json:"issuer"`
		User         UserPayload   `json:"user"`
		Client       ClientPayload `json:"client"`
	}

	GenerateAccessRefreshTokenResponse struct {
		JwtID        string        `json:"jwt_id"`
		UserName     string        `json:"user_name"`
		UserEmail    string        `json:"user_email"`
		TokenType    string        `json:"token_type"`
		Audience     string        `json:"audience"`
		Subject      string        `json:"subject"`
		Issuer       string        `json:"issuer"`
		Client       ClientPayload `json:"client"`
		RefreshToken TokenPayload  `json:"refresh_token"`
		AccessToken  TokenPayload  `json:"access_token"`
	}
)
