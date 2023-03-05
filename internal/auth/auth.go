package auth

import (
	"net/http"

	"github.com/spongeling/admin-api/internal/dao"
)

var Users []*dao.User

// Authenticator is a middleware for authentication using BasicAuth
func Authenticator(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "invalid credentials", http.StatusInternalServerError)
			return
		}

		for _, user := range Users {
			if user.Username == username && user.Password == password {
				next.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
	})

}
