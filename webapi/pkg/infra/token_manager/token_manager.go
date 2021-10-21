package token_manager

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"

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
		return "", err
	}

	privateKey, err := parseRSAPrivateKey(privateKeyInBytes)
	if err != nil {
		pst.logger.Error(err.Error())
		return "", err
	}

	claims := jwt.RegisteredClaims{
		Audience: jwt.ClaimStrings{
			tokenData.Audience,
		},
		Issuer:    os.Getenv("APP_ISSUER"),
		ExpiresAt: jwt.NewNumericDate(tokenData.ExpireIn),
		ID:        fmt.Sprintf("%d", tokenData.Id),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}

	token, err := claimsGenerator(jwt.SigningMethodRS256, claims).SignedString(privateKey)
	if err != nil {
		pst.logger.Error(err.Error())
		return "", err
	}

	return token, nil
}

func (pst tokenManager) VerifyToken(accessToken string) (*dtos.AuthenticatedUserDto, error) {
	publicKeyInBytes, err := fileReader(os.Getenv("RSA_PUBLIC_KEY_DIR"))
	if err != nil {
		pst.logger.Error(err.Error())
		return nil, err
	}

	publicKey, err := parseRSAPublicKey(publicKeyInBytes)
	if err != nil {
		pst.logger.Error(err.Error())
		return nil, err
	}

	token, err := parseClaims(
		accessToken,
		&jwt.RegisteredClaims{},
		func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
				pst.logger.Error(err.Error())
				return nil, errors.New("unexpected method")
			}
			return publicKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.ExpiresAt.Unix() < time.Now().UTC().Unix() {
		return nil, errors.New("jwt is expired")
	}

	id, _ := strconv.Atoi(claims.ID)

	return &dtos.AuthenticatedUserDto{
		Id:          id,
		AccessToken: accessToken,
		Kind:        os.Getenv("TOKEN_KIND"),
		ExpireIn:    claims.ExpiresAt.Time,
	}, nil
}

func NewTokenManager(logger interfaces.ILogger) interfaces.ITokenManager {
	return tokenManager{
		logger: logger,
	}
}
