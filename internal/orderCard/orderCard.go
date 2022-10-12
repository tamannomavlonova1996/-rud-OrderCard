package orderCard

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/models"
	"log"
)

type OrderCard models.OrderCard

func (c *OrderCard) CreateOrderCards() error {
	err := db.DB.Table("orderCard").Create(&c).Error
	if err != nil {
		log.Println("db,CreateCards err", err)
		return err
	}
	return nil
}

func (c *OrderCard) GetOrderCards() (cards []*OrderCard, err error) {
	err = db.DB.Table("orderCard").Select("*").Find(&cards).Error
	if err != nil {
		log.Println("db,GetCards err", err)
		return nil, err
	}
	return
}

func (c *OrderCard) GetOrderCardByID(id string) (*OrderCard, error) {
	err := db.DB.Table("orderCard").Where("id=?", id).First(&c).Error
	if err != nil {
		log.Println("db,GetCardByID err", err)
		return nil, err
	}
	return c, nil
}

func (c *OrderCard) UpdateOrderCardByID() error {
	err := db.DB.Table("orderCard").Where("id=?", c.ID).Update(c).Error
	if err != nil {
		log.Println("db,UpdateCardByID err", err)
		return err
	}
	return nil
}

func (c *OrderCard) DeleteOrderCardByID(id string) error {
	err := db.DB.Table("orderCard").Delete(&c, "id=?", id).Error
	if err != nil {
		log.Println("db,DeleteCardByID err", err)
		return err
	}
	return nil
}
