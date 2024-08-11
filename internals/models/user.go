package models

import (
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	GoogleID   *string   `json:"google_id,omitempty"`
}

