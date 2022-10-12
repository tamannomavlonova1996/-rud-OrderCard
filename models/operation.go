package models

import "time"

type Operation struct {
	ID                string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	SenderAccountID   string    `json:"sender_account_id" gorm:"sender_account_id"`
	ReceivedAccountID string    `json:"received_account_id" gorm:"received_account_id"`
	Sum               int64     `json:"sum" gorm:"sum"`
	Status            string    `json:"status" gorm:"status"`
	CreatedAt         time.Time `json:"created"`
	Account           Account   `gorm:"references:ID"`
}
