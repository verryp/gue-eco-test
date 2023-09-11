package presentation

type (
	ValidateTokenResponse struct {
		Client ClientPayload      `json:"client"`
		User   *AuthenticatedUser `json:"user,omitempty"`
	}
)
