package handlers

import (
	"awesomeProject2/helpers"
	user2 "awesomeProject2/internal/repository/user"
	"awesomeProject2/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"
	"time"
)

const authorizationHeader = "Authorization"

func AuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get(authorizationHeader)
		if auth == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		bearerToken := strings.Split(auth, " ")
		if bearerToken[0] != "Bearer" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		if len(bearerToken) != 2 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		_, err := ParseToken(bearerToken[1])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
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

	if helpers.CheckPasswordHash(password, userOld.Password) == false {
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
