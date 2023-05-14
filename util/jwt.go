package util

import (
	"dbo-test/config"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type authClaims struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(name string, id int64) string {
	claims := &authClaims{
		name,
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Configure.Jwt.Key))
	if err != nil {
		log.Fatal(err.Error())
	}
	return t
}
