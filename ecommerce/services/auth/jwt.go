package auth

import (
	"context"
	"ecommerce/config"
	"ecommerce/types"
	"ecommerce/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "userID"

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenAsString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenAsString, nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			log.Println("no auth header")
			permissionDenied(w)
			return
		}

		parts := strings.Split(authorization, "Bearer ")
		if len(parts) != 2 {
			log.Println("invalid authorization header")
			permissionDenied(w)
			return
		}

		authToken := parts[1]

		token, err := validateToken(authToken)

		if err != nil {
			log.Printf("failed to validate token %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		str, ok := claims["userID"].(string)

		if !ok {
			log.Println("userID not found in token")
			permissionDenied(w)
			return
		}

		userID, _ := strconv.Atoi(str)

		user, err := store.GetUserById(userID)

		if err != nil {
			log.Printf("error getting user by id %v, error: %v", userID, err)
			permissionDenied(w)
			return
		}

		if user == nil {
			log.Printf("failed to get user by id %v", userID)
			permissionDenied(w)
			return
		}

		ctx := r.Context()

		r = r.WithContext(context.WithValue(ctx, UserKey, user.ID))

		handlerFunc(w, r)
	}
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)

	if !ok {
		return -1
	}

	return userID
}
