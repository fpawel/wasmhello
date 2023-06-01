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
	UserPass
	jwt.StandardClaims
}

func JwtParse(tokenString string) (UserPass, error) {

	claims := new(accClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return UserPass{}, err
	}
	if !token.Valid {
		return UserPass{}, ErrUnauthorized
	}
	return claims.UserPass, nil
}

func JwtTokenize(accPass UserPass) (string, error) {
	claims := &accClaims{
		accPass,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
