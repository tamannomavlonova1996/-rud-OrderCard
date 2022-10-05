package models

import (
	"time"
)

type Cards struct {
	ID          string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	CompanyName string    `json:"company_name" gorm:"company_name"`
	FullName    string    `json:"full_name" gorm:"full_name"`
	Phone       string    `json:"phone" gorm:"phone"`
	Location    string    `json:"location" gorm:"location"`
	CreatedAt   time.Time `json:"created"`
}
