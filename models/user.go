package models

import "time"

type User struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	FullName  string    `json:"full_name" gorm:"full_name"`
	Email     string    `json:"email" gorm:"email"`
	Password  string    `json:"-" gorm:"password"`
	Role      string    `json:"-" gorm:"role"`
	CreatedAt time.Time `json:"created"`
}

type UserLogin struct {
	Email    string `json:"email" gorm:"email"`
	Password string `json:"-" gorm:"password"`
}
