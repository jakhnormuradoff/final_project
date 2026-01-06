package api

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if appConfig.Password != "" {
			var jwtStr string
			cookie, err := r.Cookie("token")
			if err == nil {
				jwtStr = cookie.Value
			}
			w.Header().Set("Content-Type", "application/json")
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "Authentification required"})
				return
			}

			claims := jwt.MapClaims{}

			token, err := jwt.ParseWithClaims(jwtStr, &claims, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid

				}
				return []byte(appConfig.Password), nil

			})

			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "Authentification required"})
				return
			}

		}
		next(w, r)
	})
}
