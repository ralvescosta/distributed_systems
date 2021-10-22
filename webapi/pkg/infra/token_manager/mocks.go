package token_manager

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtMocked struct{}

func (m JwtMocked) FileReader(failure bool) func(path string) ([]byte, error) {
	return func(path string) ([]byte, error) {
		if failure {
			return nil, errors.New("error when try to read rsa private key")
		}

		return []byte("key"), nil
	}
}

func (m JwtMocked) ParseRSAPrivateKey(failure bool) func(key []byte) (*rsa.PrivateKey, error) {
	return func(key []byte) (*rsa.PrivateKey, error) {
		if failure {
			return nil, errors.New("parse rsa private key")
		}

		return rsa.GenerateKey(rand.Reader, 2048)
	}

}

func (m JwtMocked) ParseRSAPublicKey(failure bool) func(key []byte) (*rsa.PublicKey, error) {
	return func(key []byte) (*rsa.PublicKey, error) {
		if failure {
			return nil, errors.New("parse rsa public key")
		}

		pKey, err := rsa.GenerateKey(rand.Reader, 2048)

		return &pKey.PublicKey, err
	}

}

func (m JwtMocked) ClaimsGenerator(failure bool) func(method jwt.SigningMethod, claims jwt.Claims) *jwt.Token {
	return func(method jwt.SigningMethod, claims jwt.Claims) *jwt.Token {
		if failure {
			return nil
		}

		return &jwt.Token{
			Header: map[string]interface{}{},
			Claims: jwt.MapClaims{},
			Method: method,
		}
	}
}

func (m JwtMocked) ParseClaims(failure bool, unixTime time.Time) func(tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc) (*jwt.Token, error) {
	return func(tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc) (*jwt.Token, error) {
		if failure {
			return nil, errors.New("")
		}

		return &jwt.Token{
			Claims: &jwt.RegisteredClaims{
				ID:        "1",
				ExpiresAt: jwt.NewNumericDate(unixTime),
			},
			Valid: true,
		}, nil
	}
}
