package product

type (
	BaseResponse struct {
		Status  string `json:"status"`
		Message string `json:"message,omitempty"`
	}

	GetProductResponse struct {
		BaseResponse
		Data *GetProduct `json:"data"`
	}

	GetProduct struct {
		ID             string  `json:"id"`
		Name           string  `json:"name"`
		QuotaPerDays   int     `json:"quota_per_days"`
		QuotaRemaining int     `json:"quota_remaining"`
		Quantity       int     `json:"quantity"`
		Category       string  `json:"category"`
		IsAvailable    string  `json:"is_available"`
		Price          float64 `json:"price"`
		CreatedAt      string  `json:"created_at"`
		UpdatedAt      string  `json:"updated_at"`
	}

	UpdateProductRequest struct {
		Name         string  `json:"name,omitempty"`
		QuotaPerDays int     `json:"quota_per_days,omitempty"`
		GrantType    string  `json:"grant_type,omitempty"`
		Quantity     int     `json:"quantity,omitempty"`
		Category     string  `json:"category,omitempty"`
		Price        float64 `json:"price,omitempty"`
	}
)
