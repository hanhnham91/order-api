package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Define a sample custom claims struct.
type CustomClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type IToken interface {
	EncodeToken(
		myClaims CustomClaims,
		secretKey string,
		expireTime time.Time,
	) (string, error)
	DecodeToken(token string, secretKey string) (CustomClaims, error)
}

type token struct{}

func NewJwt() IToken {
	return &token{}
}

func (t token) EncodeToken(
	myClaims CustomClaims,
	secretKey string,
	expireTime time.Time,
) (string, error) {
	myClaims.ExpiresAt = jwt.NewNumericDate(expireTime)
	myClaims.IssuedAt = jwt.NewNumericDate(time.Now())

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims).SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t token) DecodeToken(token string, secretKey string) (CustomClaims, error) {
	claims := CustomClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	return claims, err
}
