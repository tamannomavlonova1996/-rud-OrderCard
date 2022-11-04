package user

import (
	"awesomeProject2/helpers"
	"awesomeProject2/internal/repository/user"
	"awesomeProject2/internal/service/valid"
	"awesomeProject2/models"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"log"
)

var (
	SentPassToEmailTemplate = fmt.Sprint("Ваш пароль %s, используйте его для входа в приложение")
)

func CreateUser(req models.User) (err error) {

	err = validation.Errors{
		"id":        validation.Validate(req.ID, validation.Required, is.UUID),
		"full_name": validation.Validate(req.FullName, validation.Required, validation.Length(5, 30)),
		"email":     validation.Validate(req.Email, validation.Required, is.Email),
		"role":      validation.Validate(req.Role, validation.Required, validation.In("user", "admin")),
	}

	password := helpers.RandStrongPassword()
	checkPassword := valid.ValidationPassword(password)
	if !checkPassword {
		return fmt.Errorf("пароль не прошел валидацию: %w", err)
	}

	req.Password, err = helpers.HashPassword(password)
	if err != nil {
		log.Println("HashPassword: пароль не хешировался")
		return err
	}
	req.Role = "user"
	var user = user.User(req)
	err = user.CreateUser()
	if err != nil {
		return fmt.Errorf("не получилось создать юзера: %w", err)
	}
	msg := []byte(fmt.Sprintf(SentPassToEmailTemplate, password))
	emails := []string{user.Email}

	err = helpers.SendMessageByEmail(emails, "Ваш пароль", msg)
	if err != nil {
		return fmt.Errorf("не получилось отправить пароль на почту: %w", err)
	}

	return nil
}

func GetUsers() (users []*user.User, err error) {
	var user user.User
	users, err = user.GetUsers()
	if err != nil {
		return nil, fmt.Errorf("не получилось получить юзеров : %w", err)
	}
	return users, nil
}

func GetUserByID(id string) (user *user.User, err error) {
	err = validation.Errors{
		"id": validation.Validate(id, validation.Required, is.UUID),
	}
	user, err = user.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("не получилось получить юзера с таким айди : %w", err)
	}
	return user, nil
}

func UpdateUserByID(req models.User) (err error) {
	err = validation.Errors{
		"id":        validation.Validate(req.ID, validation.Required, is.UUID),
		"full_name": validation.Validate(req.FullName, validation.Required, validation.Length(5, 30)),
		"email":     validation.Validate(req.Email, validation.Required, is.Email),
		"role":      validation.Validate(req.Role, validation.Required, validation.In("user", "admin")),
	}
	var user = user.User(req)
	err = user.UpdateUserByID()
	if err != nil {
		return fmt.Errorf("не получилось изменить юзера: %w", err)
	}
	return nil
}

func DeleteUserByID(id string) (err error) {
	err = validation.Errors{
		"id": validation.Validate(id, validation.Required, is.UUID),
	}
	var user user.User
	err = user.DeleteUserByID(id)
	if err != nil {
		return fmt.Errorf("не получилось удалить юзера: %w", err)
	}
	return nil
}

func ResetPassword(req *models.ResetPassword) (err error) {

	var user user.User
	u, err := user.GetUserByEmail(req.Email)
	if err != nil {
		return fmt.Errorf("не получилось получить юзера по мейлу : %w", err)
	}

	password := helpers.RandStrongPassword()

	u.Password, err = helpers.HashPassword(password)
	if err != nil {
		return fmt.Errorf("не получилось хешировать пароль: %w", err)
	}

	err = u.Update()
	if err != nil {
		return fmt.Errorf("не получилось обновить пароль: %w", err)
	}

	msg := []byte(fmt.Sprintf(SentPassToEmailTemplate, password))
	emails := []string{u.Email}

	err = helpers.SendMessageByEmail(emails, "ваш пароль", msg)
	if err != nil {
		return fmt.Errorf("не получилось отправить пароль: %w", err)
	}

	return nil
}
