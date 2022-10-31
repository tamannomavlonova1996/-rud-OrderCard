package user

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/models"
	"log"
)

type User models.User

func (u *User) CreateUser() error {
	err := db.DB.Table("users").Create(&u).Error
	if err != nil {
		log.Println("db, CreateUser, err ", err)
		return err
	}
	return nil
}

func (u *User) GetUsers() (users []*User, err error) {
	err = db.DB.Table("users").Select("*").Find(&users).Error
	if err != nil {
		log.Println("db, GetUsers, err ", err)
		return
	}
	return
}

func (u *User) GetUserByID(id string) (*User, error) {
	err := db.DB.Table("users").Where("id=?", id).First(&u).Error
	if err != nil {
		log.Println("db,GetUserByID err", err)
		return nil, err
	}
	return u, nil
}

func (u *User) UpdateUserByID() error {
	err := db.DB.Table("users").Where("id=?", u.ID).Updates(User{FullName: u.FullName, Email: u.Email}).Error
	if err != nil {
		log.Println("db, UpdateUserByID err", err)
		return err
	}
	return nil
}

func (u *User) DeleteUserByID(id string) error {
	err := db.DB.Table("users").Delete(&u, "id=?", id).Error
	if err != nil {
		log.Println("db,DeleteUserByID err", err)
		return err
	}
	return nil
}

func (u *User) GetUserByEmail(email string) (*User, error) {
	err := db.DB.Table("users").Where("email=?", email).First(&u).Error
	if err != nil {
		log.Println("db,GetUserByEmail err", err)
		return nil, err
	}
	return u, nil
}
