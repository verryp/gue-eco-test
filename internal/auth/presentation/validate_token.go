package presentation

type (
	ValidateTokenResponse struct {
		Client ClientPayload      `json:"client"`
		User   *AuthenticatedUser `json:"user,omitempty"`
	}

	ValidateTokenRequest struct {
		IPAddress string `json:"ip_address"`
		UserAgent string `json:"user_agent"`
		Path      string `json:"path"`
		Token     string `json:"token"`
	}
)
