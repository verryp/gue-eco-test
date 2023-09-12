package model

import (
	"time"
)

type ActivityLog struct {
	ID           uint      `json:"id" db:"id, primarykey"`
	UserID       int64     `json:"user_id" db:"user_id"`
	IPAddress    string    `json:"ip_address" db:"ip_address"`
	UserAgent    string    `json:"user_agent" db:"user_agent"`
	PathEndpoint string    `json:"path_endpoint" db:"path_endpoint"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
