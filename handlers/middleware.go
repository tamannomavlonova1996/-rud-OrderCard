package handlers

import (
	"awesomeProject2/helpers"
	"context"
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
		userID, err := helpers.ParseToken(bearerToken[1])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := setUserID(r.Context(), userID)
		r.WithContext(ctx)

		next.ServeHTTP(w, r)
		return
	})
}

const contextUserIDKey = "user_id"

func setUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, contextUserIDKey, userID)
}

func getUserID(ctx context.Context) string {
	return ctx.Value(contextUserIDKey).(string)
}
