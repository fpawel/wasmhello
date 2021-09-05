package datatype

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	// the JWT key used to create the signature
	jwtKey = []byte("my_secret_key")
)

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type accClaims struct {
	AccPass
	jwt.StandardClaims
}

func JwtParse(tokenString string) (AccPass, error) {

	claims := new(accClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return AccPass{}, err
	}
	if !token.Valid {
		return AccPass{}, ErrUnauthorized
	}
	return claims.AccPass, nil
}

func JwtTokenize(accPass AccPass) (string, error) {
	claims := &accClaims{
		accPass,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
