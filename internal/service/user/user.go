package service

import (
	"awesomeProject2/helpers"
	user2 "awesomeProject2/internal/user"
	"fmt"
	"log"
)

var (
	SentPassToEmailTemplate = fmt.Sprint("Ваш пароль %s, используйте его для входа в приложение")
)

func CreateUser(user user2.User) (err error) {

	password := helpers.RandStringPassword(15)

	user.Password, err = helpers.HashPassword(password)
	if err != nil {
		log.Println("HashPassword: пароль не хешировался")
		return err
	}

	user.Role = "user"
	err = user.CreateUser()
	if err != nil {
		return err
	}
	msg := []byte(fmt.Sprintf(SentPassToEmailTemplate, password))
	emails := []string{user.Email}

	err = helpers.SendMessageByEmail(emails, "Ваш пароль", msg)
	if err != nil {
		return fmt.Errorf("не получилось отправить пароль на почту: %w", err)
	}

	return nil
}
