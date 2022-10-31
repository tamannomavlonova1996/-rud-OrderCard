package helpers

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/smtp"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*0123456789"

func RandStringPassword(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func SendMessageByEmail(receivers []string, subject string, msg []byte) error {
	from := viper.GetString("email.from")
	password := viper.GetString("email.password")

	log.Println("FROM:", from)
	log.Println("PASS::", password)

	addr := "smtp.yandex.ru:587"
	host := "smtp.yandex.ru"

	auth := smtp.PlainAuth("", from, password, host)

	msgTemp := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s\r\n",
		from, receivers, subject, msg)

	msg = []byte(msgTemp)

	err := smtp.SendMail(addr, auth, from, receivers, msg)
	if err != nil {
		log.Println(err)
	}

	return err
}
