package model

import (
	"time"
)

type Client struct {
	ID          int       `json:"id" db:"id, primarykey"`
	Name        string    `json:"name" db:"name"`
	APIKey      string    `json:"api_key" db:"api_key"`
	Algorithm   string    `json:"algorithm" db:"algorithm"`
	Location    string    `json:"location" db:"location"`
	PublicCert  string    `json:"public_cert" db:"public_cert"`
	PrivateCert string    `json:"private_cert" db:"private_cert"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
