package models

import "time"

type User struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	FullName  string    `json:"full_name" gorm:"full_name"`
	Email     string    `json:"email" gorm:"email"`
	Password  string    `json:"-" gorm:"password"`
	Role      string    `json:"role" gorm:"role"`
	CreatedAt time.Time `json:"created"`
}

type UserLogin struct {
	Email    string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
}

type ResetPassword struct {
	Email string `json:"email" gorm:"email"`
}
type ChangePassword struct {
	Password        string `json:"password" gorm:"password"`
	NewPassword     string `json:"new_password" gorm:"new_password"`
	ConfirmPassword string `json:"confirm_password" gorm:"confirm_password"`
}
