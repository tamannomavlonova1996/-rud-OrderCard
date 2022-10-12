package models

import "time"

type User struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	FullName  string    `json:"full_name" gorm:"full_name"`
	Login     string    `json:"login" gorm:"login"`
	Password  string    `json:"password" gorm:"password"`
	Role      string    `json:"role" gorm:"role"`
	CreatedAt time.Time `json:"created"`
}
