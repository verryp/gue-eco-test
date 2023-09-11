package model

import (
	"time"
)

type (
	Order struct {
		ID            uint64    `json:"id" db:"id, primarykey"`
		OrderSerial   string    `json:"order_serial" db:"order_serial"`
		CustomerName  string    `json:"customer_name" db:"customer_name"`
		CustomerEmail string    `json:"customer_email" db:"customer_email"`
		Status        string    `json:"status" db:"status"`
		TotalAmount   float64   `json:"total_amount" db:"total_amount"`
		ExpiredAt     time.Time `json:"expired_at" db:"expired_at"`
		CreatedAt     time.Time `json:"created_at" db:"created_at"`
		UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	}

	ParamCreateOrder struct {
		Order        *Order
		OrderHistory *OrderHistory
		OrderDetail  *OrderDetail
	}

	ParamUpdateOrder struct {
		Order        *Order
		OrderHistory *OrderHistory
	}
)
