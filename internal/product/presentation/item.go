package presentation

type (
	ItemListResponse struct {
		ID             string  `json:"id"`
		Name           string  `json:"name"`
		QuotaRemaining int     `json:"quota_remaining"`
		Quantity       int     `json:"quantity"`
		IsAvailable    string  `json:"is_available"`
		Category       string  `json:"category"`
		Price          float64 `json:"price"`
		CreatedAt      string  `json:"created_at"`
		UpdatedAt      string  `json:"updated_at"`
	}

	ItemListResponses struct {
		Items []ItemListResponse `json:"items"`
	}
)

type (
	CreateItemRequest struct {
		Name         string  `json:"name" valid:"required"`
		QuotaPerDays int     `json:"quota_per_days" valid:"required"`
		Quantity     int     `json:"quantity" valid:"required"`
		Category     string  `json:"category" valid:"required"`
		Price        float64 `json:"price" valid:"required"`
	}

	UpdateItemRequest struct {
		Name         string  `json:"name"`
		QuotaPerDays int     `json:"quota_per_days"`
		Quantity     int     `json:"quantity"`
		GrantType    string  `json:"grant_type"` // note: it should be handled from aggregator
		Category     string  `json:"category"`
		Price        float64 `json:"price"`
	}
)
