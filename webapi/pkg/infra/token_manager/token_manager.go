package token_manager

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
)

type tokenManager struct {
	logger interfaces.ILogger
}

var fileReader = ioutil.ReadFile
var parseRSAPrivateKey = jwt.ParseRSAPrivateKeyFromPEM
var parseRSAPublicKey = jwt.ParseRSAPublicKeyFromPEM
var claimsGenerator = jwt.NewWithClaims
var parseClaims = jwt.ParseWithClaims

func (pst tokenManager) GenerateToken(tokenData dtos.TokenDataDto) (string, error) {
	privateKeyInBytes, err := fileReader(os.Getenv("RSA_PRIVATE_KEY_DIR"))
	if err != nil {
		pst.logger.Error(err.Error())
		return "", errors.New("error when try to read rsa private key")
	}

	privateKey, err := parseRSAPrivateKey(privateKeyInBytes)
	if err != nil {
		pst.logger.Error(err.Error())
		return "", errors.New("error when try to create rsa private key")
	}

	claims := jwt.StandardClaims{
		Audience:  tokenData.Audience,
		Issuer:    os.Getenv("APP_ISSUER"),
		ExpiresAt: tokenData.ExpireIn.Unix(),
		Id:        fmt.Sprintf("%d", tokenData.Id),
		IssuedAt:  time.Now().Unix(),
		NotBefore: tokenData.ExpireIn.Unix(),
	}

	token, err := claimsGenerator(jwt.SigningMethodRS256, claims).SignedString(privateKey)
	if err != nil {
		pst.logger.Error(err.Error())
		return "", errors.New("error when try to create jwt")
	}

	return token, nil
}

func (pst tokenManager) VerifyToken(token string) (*dtos.AuthenticatedUserDto, error) {
	publicKeyInBytes, err := fileReader(os.Getenv("RSA_PUBLIC_KEY_DIR"))
	if err != nil {
		pst.logger.Error(err.Error())
		return nil, errors.New("error when try to read rsa public key")
	}

	publicKey, err := parseRSAPublicKey(publicKeyInBytes)
	if err != nil {
		pst.logger.Error(err.Error())
		return nil, errors.New("error when try to create rsa public key")
	}

	tok, err := parseClaims(token, &jwt.StandardClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			pst.logger.Error(err.Error())
			return nil, errors.New("unexpected method")
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := tok.Claims.(*jwt.StandardClaims)
	if !ok || !tok.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, errors.New("jwt is expired")
	}

	return &dtos.AuthenticatedUserDto{
		AccessToken: token,
		Kind:        "Bearer",
		ExpireIn:    time.Unix(claims.ExpiresAt, 0),
	}, nil
}

func NewTokenManager(logger interfaces.ILogger) interfaces.ITokenManager {
	return tokenManager{
		logger: logger,
	}
}
