package presentation

type (
	OrderListResponse struct {
		ID            string  `json:"id"`
		Serial        string  `json:"order_serial"`
		CustomerName  string  `json:"customer_name"`
		CustomerEmail string  `json:"customer_email"`
		Status        string  `json:"status"`
		TotalAmount   float64 `json:"total_amount"`
		ExpiredAt     string  `json:"expired_at"`
		CreatedAt     string  `json:"created_at"`
		UpdatedAt     string  `json:"updated_at"`
	}

	OrderListResponses struct {
		Items []OrderListResponse `json:"items"`
	}
)
