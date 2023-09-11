package model

import (
	"time"
)

type (
	OrderHistory struct {
		ID        uint64    `json:"id" db:"id, primarykey"`
		OrderID   uint64    `json:"order_id" db:"order_id"`
		Status    string    `json:"status" db:"status"`
		Remark    string    `json:"remark" db:"remark"`
		CreatedAt time.Time `json:"created_at" db:"created_at"`
	}
)
