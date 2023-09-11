package model

import (
	"time"
)

type (
	OrderDetail struct {
		ID           uint      `json:"id" db:"id, primarykey"`
		OrderID      uint64    `json:"order_id" db:"order_id"`
		ItemID       int64     `json:"item_id" db:"item_id"`
		ItemName     string    `json:"item_name" db:"item_name"`
		ItemPrice    float64   `json:"item_price" db:"item_price"`
		Quantity     int       `json:"quantity" db:"quantity"`
		TotalAmount  float64   `json:"total_amount" db:"total_amount"`
		CustomerNote string    `json:"customer_note" db:"customer_note"`
		CreatedAt    time.Time `json:"created_at" db:"created_at"`
		UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	}
)
