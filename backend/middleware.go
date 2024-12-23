package main

import (
	"database/sql"
	"net/http"
	"strings"
)

func authMiddleware(next http.Handler, db_connect *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if !isAuthUser(db_connect, token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// если все ок, токен совпал -> отдаем
		next.ServeHTTP(w, r)
	})
}
