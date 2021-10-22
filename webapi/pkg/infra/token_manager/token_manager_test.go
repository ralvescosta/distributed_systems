package token_manager

import (
	"testing"
	"time"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/infra/logger"

	"github.com/stretchr/testify/assert"
)

var manager = NewTokenManager(logger.NewLoggerSpy())
var tokenData = dtos.TokenDataDto{
	Id:       1,
	ExpireIn: time.Now().Add(time.Hour),
	Audience: "audience",
}

func Test_Should_CreateToken_Correctly(t *testing.T) {
	jwtMock := JwtMocked{}
	fileReader = jwtMock.FileReader(false)
	parseRSAPrivateKey = jwtMock.ParseRSAPrivateKey(false)
	claimsGenerator = jwtMock.ClaimsGenerator(false)

	token, err := manager.GenerateToken(tokenData)

	assert.NotNil(t, token)
	assert.Nil(t, err)
}

func Test_Should_Return_Err_If_Some_Error_Occur_In_ReadRSAPrivateKey(t *testing.T) {
	jwtMock := JwtMocked{}
	fileReader = jwtMock.FileReader(true)

	_, err := manager.GenerateToken(tokenData)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error when try to read rsa private key")
}

func Test_Should_Return_Err_If_Some_Error_Occur_In_ParseRSAPrivateKey(t *testing.T) {
	jwtMock := JwtMocked{}
	fileReader = jwtMock.FileReader(false)
	parseRSAPrivateKey = jwtMock.ParseRSAPrivateKey(true)

	_, err := manager.GenerateToken(tokenData)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "parse rsa private key")
}

// func TestCreateToken_ShouldReturnErrIfSomeErrorOccurInClaimGenerator(t *testing.T) {
// 	jwtMock := mocks.NewJwtMock(false, false, true)
// 	fileReader = jwtMock.FileReader
// 	parseRSAPrivateKey = jwtMock.ParseRSAPrivateKey
// 	claimsGenerator = jwtMock.ClaimsGenerator

// 	_, err := manager.GenerateToken(tokenData)

// 	assert.NotNil(t, err)
// }

func Test_Should_Verify_Token_Correctly(t *testing.T) {
	jwtMock := JwtMocked{}
	fileReader = jwtMock.FileReader(false)
	parseRSAPublicKey = jwtMock.ParseRSAPublicKey(false)
	parseClaims = jwtMock.ParseClaims(false, time.Now().Add(time.Hour))

	token, err := manager.VerifyToken("token")

	assert.NotNil(t, token)
	assert.Nil(t, err)
}

func Test_Should_Return_Err_If_Some_Error_Occur_In_ReadRSAPublicKey(t *testing.T) {
	jwtMock := JwtMocked{}
	fileReader = jwtMock.FileReader(true)

	token, err := manager.VerifyToken("token")

	assert.Nil(t, token)
	assert.Equal(t, err.Error(), "error when try to read rsa private key")
}

func Test_Should_Return_Err_If_Some_Error_Occur_In_ParseRSAPublicKey(t *testing.T) {
	jwtMock := JwtMocked{}
	fileReader = jwtMock.FileReader(false)
	parseRSAPublicKey = jwtMock.ParseRSAPublicKey(true)

	_, err := manager.VerifyToken("token")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "parse rsa public key")
}

func Test_Should_Return_Err_If_Some_Error_Occur_When_ParseClaims(t *testing.T) {
	jwtMock := JwtMocked{}
	fileReader = jwtMock.FileReader(false)
	parseRSAPublicKey = jwtMock.ParseRSAPublicKey(false)
	parseClaims = jwtMock.ParseClaims(true, time.Now().Add(time.Hour))

	_, err := manager.VerifyToken("token")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "")
}

func Test_Should_Return_Err_If_Token_Expired(t *testing.T) {
	jwtMock := JwtMocked{}
	fileReader = jwtMock.FileReader(false)
	parseRSAPublicKey = jwtMock.ParseRSAPublicKey(false)
	parseClaims = jwtMock.ParseClaims(false, time.Now().Add(-time.Hour))

	_, err := manager.VerifyToken("token")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "jwt is expired")
}
