package presentation

type (
	CheckoutRequest struct {
		CartID string `json:"cart_id" valid:"required"`
	}

	CheckoutResponse struct {
		ID            string              `json:"id"`
		Serial        string              `json:"order_serial"`
		CustomerName  string              `json:"customer_name"`
		CustomerEmail string              `json:"customer_email"`
		Status        string              `json:"status"`
		TotalAmount   float64             `json:"total_amount"`
		ExpiredAt     string              `json:"expired_at"`
		CreatedAt     string              `json:"created_at"`
		UpdatedAt     string              `json:"updated_at"`
		Detail        OrderDetailResponse `json:"detail,omitempty"`
	}

	OrderDetailResponse struct {
		ID           int     `json:"id"`
		ItemID       string  `json:"item_id"`
		ItemName     string  `json:"item_name"`
		ItemPrice    float64 `json:"item_price"`
		Quantity     int     `json:"quantity"`
		TotalAmount  float64 `json:"total_amount"`
		CustomerNote string  `json:"customer_note"`
	}
)
