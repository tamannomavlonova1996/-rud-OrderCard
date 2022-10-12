package models

import (
	"time"
)

type OrderCard struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	UserID    string    `json:"user_id" gorm:"user_id"`
	Status    string    `json:"status" gorm:"status"`
	CreatedAt time.Time `json:"created"`
	User      User      `gorm:"references:ID"`
}
