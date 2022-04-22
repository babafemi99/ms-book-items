package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"time"
)

type SignedDetails struct {
	UserID int64
	jwt.StandardClaims
}

var secretKey = []byte("NOMORESAPA")

func ValidateToken(request *http.Request) (claims *SignedDetails, msg string) {
	extractedToken, extractErr := ExtractToken(request)
	if extractErr != nil {
		fmt.Println(extractErr)
		return nil, ""
	}
	SignedToken := extractedToken
	token, err := jwt.ParseWithClaims(
		SignedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "Invalid Token"
		msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		msg = err.Error()
		return
	}
	return claims, msg
}
func ExtractToken(request *http.Request) (string, error) {
	bearerToken := request.Header.Get("Authorization")
	tokenStr := strings.Split(bearerToken, " ")
	if len(tokenStr) != 2 {
		return "", errors.New("invalid token")
	}
	return tokenStr[1], nil
}
