package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SingIn struct {
	Password string `json:"password"`
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}
	singIn := SingIn{}
	err := json.NewDecoder(r.Body).Decode(&singIn)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to decode request body"})
		return
	}

	serverPassword := appConfig.Password

	if serverPassword == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Password is not set on server"})
		return

	}

	if singIn.Password != serverPassword {
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid password"})
		return
	}
	claims := jwt.MapClaims{
		"authorized": true,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString([]byte(appConfig.Password))

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to generate token"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   signedToken,
		Expires: time.Now().Add(8 * time.Hour),
	})

	json.NewEncoder(w).Encode(map[string]string{"token": signedToken})
}
