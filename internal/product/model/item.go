package model

import (
	"time"
)

type Item struct {
	ID           uint64    `json:"id" db:"id, primarykey"`
	Name         string    `json:"name" db:"name"`
	QuotaPerDays int       `json:"quota_per_days" db:"quota_per_days"`
	Quantity     int       `json:"quantity" db:"quantity"`
	Category     string    `json:"category" db:"category"`
	Price        float64   `json:"price" db:"price"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
