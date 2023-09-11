package presentation

type (
	ReTokenRequest struct {
		Token    string `json:"token"`
		ClientID string `json:"client_id"`
		PathURL  string `json:"path_url"`
	}
)
