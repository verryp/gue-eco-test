package presentation

type (
	SignInRequest struct {
		ClientID string `json:"client_id" valid:"required"`
		Email    string `json:"email" valid:"required,email"`
		Password string `json:"password" valid:"required"`
	}

	SignInResponse struct {
		TokenType    string       `json:"token_type"`
		AccessToken  TokenPayload `json:"access_token"`
		RefreshToken TokenPayload `json:"refresh_token"`
	}

	TokenPayload struct {
		Value     string `json:"value"`
		ExpiredAt int64  `json:"expired_at"`
	}
)
