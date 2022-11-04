package card

import (
	"awesomeProject2/internal/repository/card"
	"awesomeProject2/models"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func CreateCard(req models.Card) (err error) {
	err = validation.Errors{
		"user_id": validation.Validate(req.UserID, validation.Required, is.UUID),
		"pan":     validation.Validate(req.PAN, validation.Required, is.Digit, validation.Length(18, 18)),
		"period":  validation.Validate(req.Period, validation.Required, validation.Date("2022-10-01")),
		"cvv":     validation.Validate(req.CVV, validation.Required, is.Digit, validation.Length(3, 3)),
	}
	req.Status = "в обработке"
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
	err = validation.Errors{
		"id": validation.Validate(id, validation.Required, is.UUID),
	}
	card, err = card.GetCardByID(id)
	if err != nil {
		return nil, fmt.Errorf("не получилосӣ найти дннҷе по карту по такому айди: %w", err)
	}
	return card, nil
}

func UpdateCardByID(req models.Card) (err error) {
	err = validation.Errors{
		"user_id": validation.Validate(req.UserID, validation.Required, is.UUID),
		"pan":     validation.Validate(req.PAN, validation.Required, is.Digit, validation.Length(18, 18)),
		"period":  validation.Validate(req.Period, validation.Required, validation.Date("2022-10-01")),
		"cvv":     validation.Validate(req.CVV, validation.Required, is.Digit, validation.Length(3, 3)),
		"status":  validation.Validate(req.Status, validation.Required, validation.In("в обработке", "одобрено", "отказано")),
	}
	var card = card.Card(req)
	err = card.UpdateCardByID()
	if err != nil {
		return fmt.Errorf("не получилось изменить данные карты: %w", err)
	}
	return nil
}

func DeleteCardByID(id string) (err error) {
	err = validation.Errors{
		"user_id": validation.Validate(id, validation.Required, is.UUID),
	}
	var card card.Card
	err = card.DeleteCardByID(id)
	if err != nil {
		return fmt.Errorf("не получилось удалить карту %w", err)
	}

	return nil
}
