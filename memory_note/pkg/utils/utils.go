package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JWTsecret = []byte("AC")

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_Name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string, password string) (string, error) {
	notTime := time.Now()
	expireTime := notTime.Add(120 * time.Minute)
	claims := Claims{
		Id:       id,
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "list",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := tokenClaims.SignedString(JWTsecret)
	return token, err
}
func ParseToken(token string) (*Claims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaim != nil {
		if claims, ok := tokenClaim.Claims.(*Claims); ok && tokenClaim.Valid {
			return claims, nil
		}
	}
	return nil, err
}
