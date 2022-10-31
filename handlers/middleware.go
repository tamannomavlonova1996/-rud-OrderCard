package handlers

import (
	"awesomeProject2/helpers"
	"net/http"
	"strings"
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
		_, err := helpers.ParseToken(bearerToken[1])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
}
