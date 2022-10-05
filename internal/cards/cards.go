package cards

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/models"
	"log"
)

type Cards models.Cards

func (c *Cards) CreateCards() error {
	err := db.DB.Table("cards").Create(&c).Error
	if err != nil {
		log.Println("db,CreateCards err", err)
		return err
	}
	return nil
}

func (c *Cards) GetCards() (cards []*Cards, err error) {
	err = db.DB.Table("cards").Select("*").Find(&cards).Error
	if err != nil {
		log.Println("db,GetCards err", err)
		return nil, err
	}
	return
}

func (c *Cards) GetCardByID(id string) (*Cards, error) {
	err := db.DB.Table("cards").Where("id=?", id).First(&c).Error
	if err != nil {
		log.Println("db,GetCards err", err)
		return nil, err
	}
	return c, nil
}

func (c *Cards) UpdateCardByID() error {
	err := db.DB.Table("cards").Where("id=?", c.ID).Update(c).Error
	if err != nil {
		log.Println("db,GetCards err", err)
		return err
	}
	return nil
}

func (c *Cards) DeleteCardByID(id string) error {
	err := db.DB.Table("cards").Delete(&c, "id=?", id).Error
	if err != nil {
		log.Println("db,GetCards err", err)
		return err
	}
	return nil
}
