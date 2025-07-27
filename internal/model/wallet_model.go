package model

import "time"

type Wallet struct {
	ID        string    `gorm:"primaryKey;autoIncrement"`
	UserId    string    `json:"user_id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
