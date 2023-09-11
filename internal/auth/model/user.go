package model

import (
	"time"
)

type (
	User struct {
		ID        uint64    `json:"id" db:"id, primarykey"`
		Name      string    `json:"name" db:"name"`
		Email     string    `json:"email" db:"email"`
		Password  string    `json:"password" db:"password"`
		CreatedAt time.Time `json:"created_at" db:"created_at"`
		UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	}
)
