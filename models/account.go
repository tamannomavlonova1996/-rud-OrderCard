package models

import "time"

type Account struct {
	ID            string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	CardID        string    `json:"card_id" gorm:"card_id"`
	AccountNumber string    `json:"account_number" gorm:"column:account_number;size:18"`
	Currency      string    `json:"currency" gorm:"currency"`
	Balance       string    `json:"balance" gorm:"balance"`
	CreatedAt     time.Time `json:"created"`
	Card          Card      `gorm:"references:ID"`
}
