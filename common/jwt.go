package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var Key = "zff_test"

type OceanClaims struct {
	Username string `json:"username"`
	Uid      uint   `json:"id"`
	jwt.StandardClaims
}

//获取token
func GenToken(username string, uid uint) string {
	claims := OceanClaims{
		Username: username,
		Uid:      uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(Key))
	if err != nil {
		//TODO
	}
	return tokenString
}

//解析token
func ParseToken(tokenString string) (*jwt.Token, OceanClaims, error) {
	claims := OceanClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})
	return token, claims, err

}
