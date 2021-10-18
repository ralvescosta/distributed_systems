package hasher

import (
	"golang.org/x/crypto/bcrypt"

	"webapi/pkg/app/errors"
	"webapi/pkg/app/interfaces"
)

type hasher struct {
	logger interfaces.ILogger
}

var generateHash = bcrypt.GenerateFromPassword
var compareHash = bcrypt.CompareHashAndPassword

func (pst hasher) Hahser(text string) (string, error) {
	hashed, err := generateHash([]byte(text), 9)
	if err != nil {
		pst.logger.Error(err.Error())
		return "", errors.NewInternalError(err.Error())
	}

	return string(hashed), nil
}

func (h hasher) Verify(originalText, hashedText string) bool {
	if err := compareHash([]byte(hashedText), []byte(originalText)); err != nil {
		return false
	}

	return true
}

func NewHahser(logger interfaces.ILogger) interfaces.IHasher {
	return &hasher{
		logger: logger,
	}
}
