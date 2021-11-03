package hasher

import (
	"errors"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/infra/logger"
)

type bcryptMocked struct {
	failure bool
	hash    []byte
}

func (m bcryptMocked) GenerateHash(password []byte, cost int) ([]byte, error) {
	if m.failure {
		return []byte(""), errors.New("bcrypt error")
	}

	return m.hash, nil
}

func (m bcryptMocked) CompareHash(hashedPassword, password []byte) error {
	if m.failure {
		return errors.New("bcrypt error")
	}

	return nil
}

func newBcryptMock(failure bool, hash []byte) *bcryptMocked {
	return &bcryptMocked{
		failure: failure,
		hash:    hash,
	}
}

type hasherToTest struct {
	hasher    interfaces.IHasher
	loggerSpy interfaces.ILogger
}

func newHasherToTest() hasherToTest {
	loggerSpy := logger.NewLoggerSpy()

	return hasherToTest{
		hasher:    NewHahser(loggerSpy),
		loggerSpy: loggerSpy,
	}
}
