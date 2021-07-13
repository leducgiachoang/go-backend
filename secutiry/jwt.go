package secutiry

import (
	"github.com/dgrijalva/jwt-go"
	"go-backend/model"
	"os"
	"strconv"
	"time"
)

const JwtKey = "9420ecb41a44b48e9cd1c8019997498ed"

func GenToken(user model.User) (string, error) {
	jwtExpConfig := os.Getenv("jwtExpires")
	jwtExpValue, _ := strconv.Atoi(jwtExpConfig)

	jwrExpDuration := time.Hour * time.Duration(jwtExpValue)

	claims := &model.JwtCustomClaims{
		UserId: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwrExpDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
