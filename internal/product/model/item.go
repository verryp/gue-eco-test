package model

import (
	"time"
)

type (
	Item struct {
		ID           uint64    `json:"id" db:"id, primarykey"`
		Name         string    `json:"name" db:"name"`
		QuotaPerDays int       `json:"quota_per_days" db:"quota_per_days"`
		Quantity     int       `json:"quantity" db:"quantity"`
		Category     string    `json:"category" db:"category"`
		Price        float64   `json:"price" db:"price"`
		CreatedAt    time.Time `json:"created_at" db:"created_at"`
		UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	}

	Quota struct {
		ID     uint  `json:"id" db:"id, primarykey"`
		ItemID int64 `json:"item_id" db:"item_id"`

		// DateLimiter will use for scheduler for new the remaining quota and comparer
		DateLimiter time.Time `json:"date_limiter" db:"date_limiter"`

		QuotaRemaining int       `json:"quota_remaining" db:"quota_remaining"`
		CreatedAt      time.Time `json:"created_at" db:"created_at"`
		UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	}

	ParamCreateItem struct {
		Item  *Item
		Quota *Quota
	}

	ItemQuota struct {
		ID           uint64 `json:"id" db:"id, primarykey"`
		Name         string `json:"name" db:"name"`
		QuotaPerDays int    `json:"quota_per_days" db:"quota_per_days"`

		// DateLimiter scheduler for new the remaining quota and comparer
		DateLimiter time.Time `json:"date_limiter" db:"date_limiter"`

		QuotaRemaining int       `json:"quota_remaining" db:"quota_remaining"`
		Quantity       int       `json:"quantity" db:"quantity"`
		Category       string    `json:"category" db:"category"`
		Price          float64   `json:"price" db:"price"`
		CreatedAt      time.Time `json:"created_at" db:"created_at"`
		UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	}
)
