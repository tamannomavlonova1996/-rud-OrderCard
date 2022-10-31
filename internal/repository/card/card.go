package card

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/models"
	"log"
)

type Card models.Card

func (cd *Card) CreateCard() error {
	err := db.DB.Table("cards").Create(&cd).Error
	if err != nil {
		log.Println("db,CreateCards err", err)
		return err
	}
	return nil
}

func (cd *Card) GetCards() (cards []*Card, err error) {
	err = db.DB.Table("cards").
		Select("*").Preload("User").
		Find(&cards).Error
	if err != nil {
		log.Println("db,GetCards err", err)
		return nil, err
	}
	return
}

func (cd *Card) GetCardByID(id string) (*Card, error) {
	err := db.DB.Table("cards").Where("id=?", id).Preload("User").First(&cd).Error
	if err != nil {
		log.Println("db,GetCards err", err)
		return nil, err
	}
	return cd, nil
}

func (cd *Card) UpdateCardByID() error {
	err := db.DB.Table("cards").Where("id=?", cd.ID).Update(cd).Error
	if err != nil {
		log.Println("db,UpdateCardByID err", err)
		return err
	}
	return nil
}

func (cd *Card) DeleteCardByID(id string) error {
	err := db.DB.Table("cards").Delete(&cd, "id=?", id).Error
	if err != nil {
		log.Println("db,DeleteCardByID err", err)
		return err
	}
	return nil
}
