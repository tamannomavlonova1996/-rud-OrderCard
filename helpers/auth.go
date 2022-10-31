package helpers

import (
	user2 "awesomeProject2/internal/repository/user"
	"awesomeProject2/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"
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

const (
	tokenTTL = 12 * time.Hour
	salt     = "sjakfslkaf23j213123kjklkjl"
)

var (
	SigningKey = "token_password"
	numberSet  = "0123456789"
	response   = models.Response{
		Code: http.StatusOK,
	}
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func CreateToken(email string, password string) (string, error) {
	var user user2.User
	userOld, err := user.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if CheckPasswordHash(password, userOld.Password) == false {
		response.Code = http.StatusBadRequest
		log.Println("Неверные данные")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userOld.ID,
	})
	tokenString, err := token.SignedString([]byte(SigningKey))

	return tokenString, nil
}

func ParseToken(accessToken string) (string, error) {

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		return []byte(SigningKey), nil
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("Token claims are not type *tokenClaims")
	}

	return claims.UserId, nil
}
