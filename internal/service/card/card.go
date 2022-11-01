package card

import (
	"awesomeProject2/internal/repository/card"
	"awesomeProject2/models"
	"fmt"
)

func CreateCard(req models.Card) (err error) {
	cd := card.Card(req)
	err = cd.CreateCard()
	if err != nil {
		return fmt.Errorf("не удалось создать карту: %w", err)
	}
	return nil
}

func GetCards() (cards []*card.Card, err error) {
	var card card.Card
	cards, err = card.GetCards()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить данные по картам: %w", err)
	}
	return cards, nil
}

func GetCardByID(id string) (card *card.Card, err error) {
	card, err = card.GetCardByID(id)
	if err != nil {
		return nil, fmt.Errorf("не получилосӣ найти дннҷе по карту по такому айди: %w", err)
	}
	return card, nil
}

func UpdateCardByID(req models.Card) (err error) {
	var card = card.Card(req)
	err = card.UpdateCardByID()
	if err != nil {
		return fmt.Errorf("не получилось изменить данные карты: %w", err)
	}
	return nil
}

func DeleteCardByID(id string) (err error) {
	var card card.Card
	err = card.DeleteCardByID(id)
	if err != nil {
		return fmt.Errorf("не получилось удалить карту %w", err)
	}

	return nil
}
