package middleware

import (
	// "fmt"
	"net/http"
	"strings"

	"github.com/backendGo-learn/src/helpers"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		headerToken := r.Header.Get("Authorization")
		var response helpers.Response

		if !strings.Contains(headerToken, "Bearer") {
			response.ResponseJSON(401, "Invalid Header Type").Send(w)
			return
		}

		token := strings.Replace(headerToken, "Bearer ", "", -1)

		checkToken, err := helpers.CheckToken(token)
		if err != nil {
			response.ResponseJSON(401, "invalid token").Send(w)
			return
		}

		if !checkToken {
			response.ResponseJSON(401, "Please login back").Send(w)
			return
		}

		next.ServeHTTP(w, r)
	}
}
