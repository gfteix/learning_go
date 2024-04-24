package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func WithJWTAuth(handlerFunc http.HandlerFunc, store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token from the request (Auth header)
		tokenString := GetTokenFromRequest(r)
		// validate the token
		token, err := validateJWT(tokenString)

		if err != nil {
			log.Printf("failed to authenticate token")
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Printf("token is not valid")
			permissionDenied(w)
			return
		}
		// get the user id from the token
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string)

		_, err = store.GetUserById(userId)

		if err != nil {
			log.Printf("fail to get user")
			permissionDenied(w)
			return
		}
		// call the handler func and continue to the endpoint
		handlerFunc(w, r)
	}
}

func permissionDenied(w http.ResponseWriter) {
	WriteJson(w, http.StatusUnauthorized, ErrorResponse{
		Error: fmt.Errorf("permission denied").Error(),
	})
}

func GetTokenFromRequest(request *http.Request) string {
	tokenFromAuth := request.Header.Get("Authorization")
	tokenFromQuery := request.URL.Query().Get("token")

	if tokenFromAuth != "" {
		return tokenFromAuth
	}

	if tokenFromQuery != "" {
		return tokenFromQuery
	}

	return ""
}

func validateJWT(token string) (*jwt.Token, error) {
	secret := Envs.JWTSecret

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["algs"])
		}

		return []byte(secret), nil
	})
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CreateJWT(secret []byte, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(int(userId)),
		"expiresAt": time.Now().Add(time.Hour * 24 * 120).Unix(),
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
