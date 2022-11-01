package account

import (
	"awesomeProject2/internal/repository/account"
	"awesomeProject2/models"
	"fmt"
)

func CreateAccount(req models.Account) (err error) {
	acc := account.Account(req)
	err = acc.CreateAccount()
	if err != nil {
		return fmt.Errorf("не удалось создать акаунт: %w", err)
	}
	return nil
}

func GetAccounts() (accounts []*account.Account, err error) {
	var account account.Account
	accounts, err = account.GetAccounts()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить акаунты: %w", err)
	}
	return accounts, nil
}

func GetAccountByID(id string) (account *account.Account, err error) {
	account, err = account.GetAccountByID(id)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить акаунто такому айди: %w", err)
	}
	return account, nil
}

func UpdateAccountByID(req models.Account) (err error) {
	acc := account.Account(req)
	err = acc.UpdateAccountByID()
	if err != nil {
		return fmt.Errorf("не удалось обновить акаунт: %w", err)
	}
	return nil
}

func DeleteAccountByID(id string) (err error) {
	var account account.Account
	err = account.DeleteAccountByID(id)
	if err != nil {
		return fmt.Errorf("не удалось удалить акаунт: %w", err)
	}
	return nil
}
