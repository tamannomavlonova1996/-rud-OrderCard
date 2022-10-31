package operation

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/models"
	"log"
)

type Operation models.Operation

func (o *Operation) CreateOperation() error {
	err := db.DB.Table("operations").Create(&o).Error
	if err != nil {
		log.Println("db, CreateOperation, err ", err)
		return err
	}
	return nil
}

func (o *Operation) GetOperations() (operations []*Operation, err error) {
	err = db.DB.Table("operations").Select("*").Preload("Account").Find(&operations).Error
	if err != nil {
		log.Println("db, GetOperations, err ", err)
		return
	}
	return
}

func (o *Operation) GetOperationByID(id string) (*Operation, error) {
	err := db.DB.Table("operations").Where("id=?", id).Preload("Account").First(&o).Error
	if err != nil {
		log.Println("db,GetOperationByID err", err)
		return nil, err
	}
	return o, nil
}

func (o *Operation) UpdateOperationByID() error {
	err := db.DB.Table("operations").Where("id=?", o.ID).Update(o).Error
	if err != nil {
		log.Println("db, UpdateOperationByID err", err)
		return err
	}
	return nil
}

func (o *Operation) DeleteOperationByID(id string) error {
	err := db.DB.Table("operations").Delete(&o, "id=?", id).Error
	if err != nil {
		log.Println("db,DeleteOperationByID err", err)
		return err
	}
	return nil
}
