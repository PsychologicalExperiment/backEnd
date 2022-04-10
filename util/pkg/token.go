package pkg

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Email string
	jwt.StandardClaims
}

func GenerateUserToken(
	email, secretKey string,
	expireDuration time.Duration,
) (string, error) {
	expire := time.Now().Add(expireDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, UserClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	})
	return token.SignedString([]byte(secretKey))
}

func ParseUserToken(
	token, secretKey string,
) (*UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if userClaims, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
			return userClaims, nil
		}
	}
	return nil, err
}
