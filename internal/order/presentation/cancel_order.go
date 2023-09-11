package presentation

type (
	CancelOrderRequest struct {
		OrderID string `json:"order_id" valid:"required"`
		Status  string `json:"status" valid:"required"`
		Reason  string `json:"reason" valid:"required"`
	}
)
