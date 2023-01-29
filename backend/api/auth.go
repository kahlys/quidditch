package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// User struct to hold user information
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	UserID int `json:"userid"`
	jwt.RegisteredClaims
}

// Users hold all registered users (TODO: use db instead)
var Users []User

// TODO: better generation
var jwtSecret = []byte("secret")

func generateAccessToken(userID int) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&Claims{
			UserID: userID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)),
			},
		},
	)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateAccessToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, nil, err
	}
	if !token.Valid {
		return nil, nil, fmt.Errorf("invalid token")
	}

	return token, claims, nil
}

type contextKey string

var ctxUserID = contextKey("userID")

// mwAuth is a middleware for authentication
func (h handler) mwAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, claims, err := validateAccessToken(c.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Passer l'userID au handler protégé
		ctx := context.WithValue(r.Context(), ctxUserID, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
