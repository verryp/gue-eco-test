package presentation

type (
	AddCartRequest struct {
		CustomerName  string `json:"customer_name"`
		CustomerEmail string `json:"customer_email"`
		ItemID        string `json:"item_id" valid:"required"`
		Quantity      int    `json:"quantity" valid:"required"`
		CustomerNote  string `json:"customer_note,omitempty" valid:"required"`
	}
)
