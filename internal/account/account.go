package account

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/models"
	"log"
)

type Account models.Account

func (a *Account) CreateAccount() error {
	err := db.DB.Table("account").Create(&a).Error
	if err != nil {
		log.Println("db, CreateAccount, err ", err)
		return err
	}
	return nil
}

func (a *Account) GetAccounts() (accounts []*Account, err error) {
	err = db.DB.Table("account").Select("*").Find(&accounts).Error
	if err != nil {
		log.Println("db, GetAccounts, err ", err)
		return
	}
	return
}

func (a *Account) GetAccountByID(id string) (*Account, error) {
	err := db.DB.Table("account").Where("id=?", id).First(&a).Error
	if err != nil {
		log.Println("db,GetUAccountByID err", err)
		return nil, err
	}
	return a, nil
}

func (a *Account) UpdateAccountByID() error {
	err := db.DB.Table("account").Where("id=?", a.ID).Update(&a).Error
	if err != nil {
		log.Println("db, UpdateAccountByID err", err)
		return err
	}
	return nil
}

func (a *Account) DeleteAccountByID(id string) error {
	err := db.DB.Table("account").Delete(&a, "id=?", id).Error
	if err != nil {
		log.Println("db,DeleteAccountByID err", err)
		return err
	}
	return nil
}
