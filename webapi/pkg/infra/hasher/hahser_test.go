package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"webapi/pkg/app/errors"
)

func Test_Should_Execute_Hahser_Correctly(t *testing.T) {
	var bcryptMock = newBcryptMock(false, []byte("hashed"))
	generateHash = bcryptMock.GenerateHash

	sut := newHasherToTest()

	hashed, err := sut.hasher.Hahser("text")

	assert.Equal(t, err, nil)
	assert.Equal(t, hashed, "hashed")
}

func Test_Should_Execute_Hahser_When_Same_Error_Occur_In_Crypto(t *testing.T) {
	var bcryptMock = newBcryptMock(true, []byte("hash"))
	generateHash = bcryptMock.GenerateHash

	sut := newHasherToTest()

	_, err := sut.hasher.Hahser("text")

	assert.Error(t, err)
	assert.IsType(t, err, errors.InternalError{})
}

func Test_Should_Return_Ture_If_Hash_Is_Correctly(t *testing.T) {
	var bcryptMock = newBcryptMock(false, []byte(""))
	compareHash = bcryptMock.CompareHash

	sut := newHasherToTest()

	result := sut.hasher.Verify("text", "hash")

	assert.True(t, result)
}

func Test_Should_Retorne_False_If_Hash_Is_Wrong(t *testing.T) {
	var bcryptMock = newBcryptMock(true, []byte(""))
	compareHash = bcryptMock.CompareHash

	sut := newHasherToTest()

	result := sut.hasher.Verify("text", "hash")

	assert.False(t, result)
}
