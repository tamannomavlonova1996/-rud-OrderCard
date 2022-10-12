package models

import "time"

type Card struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	UserID    string    `json:"user_id" gorm:"user_id"`
	PAN       string    `json:"pan" gorm:"column:pan;size:18"`
	Period    time.Time `json:"validity" gorm:"validity"`
	CVV       int64     `json:"cvv" gorm:"cvv"`
	Status    string    `json:"status" gorm:"status"`
	CreatedAt time.Time `json:"created"`
	User      User      `gorm:"references:ID"`
}
