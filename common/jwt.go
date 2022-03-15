package common

import (
	"awesomeProject3/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwt_key = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirtionTime := time.Now().Add(7 * 24 * time.Hour)
	claim := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirtionTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(jwt_key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwt_key, nil
	})
	return token, claims, err
}
