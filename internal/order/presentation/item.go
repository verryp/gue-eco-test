package presentation

type ItemResponse struct {
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
