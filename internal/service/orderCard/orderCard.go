package orderCard

import (
	"awesomeProject2/internal/repository/orderCard"
	"awesomeProject2/models"
	"fmt"
)

func CreateOrderCard(req models.OrderCard) (err error) {
	var orderCard = orderCard.OrderCard(req)
	err = orderCard.CreateOrderCards()
	if err != nil {
		return fmt.Errorf("не получилось создать : %w", err)
	}
	return nil
}

func GetOrderCards() (orderCards []*orderCard.OrderCard, err error) {
	var orderCard orderCard.OrderCard
	orderCards, err = orderCard.GetOrderCards()
	if err != nil {
		return nil, fmt.Errorf("не получилось выгрузить : %w", err)
	}
	return orderCards, nil
}

func GetOrderCardByID(id string) (orderCard *orderCard.OrderCard, err error) {
	orderCard, err = orderCard.GetOrderCardByID(id)
	if err != nil {
		return nil, fmt.Errorf("не получилось выгрузить с такой айди заказной карту : %w", err)
	}
	return orderCard, nil
}

func UpdateOrderCardByID(req models.OrderCard) (err error) {
	var orderCard = orderCard.OrderCard(req)
	err = orderCard.UpdateOrderCardByID()
	if err != nil {
		return fmt.Errorf("не получилось обновить данные: %w", err)
	}
	return nil
}

func DeleteOrderCardByID(id string) (err error) {
	var orderCard orderCard.OrderCard
	err = orderCard.DeleteOrderCardByID(id)
	if err != nil {
		return fmt.Errorf("не получилось удалить данные: %w", err)
	}
	return nil
}
