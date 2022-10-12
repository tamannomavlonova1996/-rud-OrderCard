package card

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/models"
	"log"
)

type Card models.Card

func (cd *Card) CreateCard() error {
	err := db.DB.Table("card").Create(&cd).Error
	if err != nil {
		log.Println("db,CreateCards err", err)
		return err
	}
	return nil
}

func (cd *Card) GetCards() (cards []*Card, err error) {
	err = db.DB.Table("card").Select("*").Find(&cards).Error
	if err != nil {
		log.Println("db,GetCards err", err)
		return nil, err
	}
	return
}

func (cd *Card) GetCardByID(id string) (*Card, error) {
	err := db.DB.Table("card").Where("id=?", id).First(&cd).Error
	if err != nil {
		log.Println("db,GetCards err", err)
		return nil, err
	}
	return cd, nil
}

func (cd *Card) UpdateCardByID() error {
	err := db.DB.Table("card").Where("id=?", cd.ID).Update(cd).Error
	if err != nil {
		log.Println("db,UpdateCardByID err", err)
		return err
	}
	return nil
}

func (cd *Card) DeleteCardByID(id string) error {
	err := db.DB.Table("card").Delete(&cd, "id=?", id).Error
	if err != nil {
		log.Println("db,DeleteCardByID err", err)
		return err
	}
	return nil
}
