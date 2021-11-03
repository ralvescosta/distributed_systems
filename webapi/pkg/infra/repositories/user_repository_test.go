package repositories

import (
	"context"
	"testing"
	"webapi/pkg/domain/dtos"

	"github.com/stretchr/testify/assert"
)

func Test_Should_Create_User_And_Returns_When_Execute_Correctly(t *testing.T) {
	sut := newUserRepositoryToTest()

	query := "INSERT INTO users \\(name, email, password\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING \\*"
	rows := sut.sqlMock.NewRows(
		[]string{"id", "name", "email", "password", "created_at", "updated_at", "deleted_at"},
	).AddRow(
		sut.mockedUser.Id,
		sut.mockedUser.Name,
		sut.mockedUser.Email,
		sut.mockedUser.Password,
		sut.mockedUser.CreatedAt,
		sut.mockedUser.UpdatedAt,
		sut.mockedUser.DeletedAt,
	)

	prep := sut.sqlMock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(sut.mockedUser.Name, sut.mockedUser.Email, sut.mockedUser.Password).WillReturnRows(rows)

	result, err := sut.repo.Create(
		context.Background(),
		sut.newrelic.StartTransaction("txn"),
		dtos.CreateUserDto{
			Name:     sut.mockedUser.Name,
			Email:    sut.mockedUser.Email,
			Password: sut.mockedUser.Password,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, sut.mockedUser.Email, result.Email)
}

func Test_Should_Returns_An_Error_When_Prepare_Return_Error(t *testing.T) {
	sut := newUserRepositoryToTest()

	sut.sqlMock.ExpectPrepare("")

	_, err := sut.repo.Create(
		context.Background(),
		sut.newrelic.StartTransaction("txn"),
		dtos.CreateUserDto{
			Name:     sut.mockedUser.Name,
			Email:    sut.mockedUser.Email,
			Password: sut.mockedUser.Password,
		},
	)

	assert.Error(t, err)
}

func Test_Should_Returns_An_Error_When_Occur_Error_In_Query_Execution(t *testing.T) {
	sut := newUserRepositoryToTest()

	query := "INSERT INTO users \\(name, email, password\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING \\*"
	rows := sut.sqlMock.NewRows(
		[]string{"id", "name", "email", "password", "created_at", "updated_at", "deleted_at"},
	).AddRow(
		sut.mockedUser.Id,
		sut.mockedUser.Name,
		sut.mockedUser.Email,
		sut.mockedUser.Password,
		sut.mockedUser.CreatedAt.String(),
		sut.mockedUser.UpdatedAt.String(),
		sut.mockedUser.DeletedAt,
	)

	prep := sut.sqlMock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(sut.mockedUser.Name, sut.mockedUser.Email, sut.mockedUser.Password).WillReturnRows(rows)

	_, err := sut.repo.Create(
		context.Background(),
		sut.newrelic.StartTransaction("txn"),
		dtos.CreateUserDto{
			Name:     sut.mockedUser.Name,
			Email:    sut.mockedUser.Email,
			Password: sut.mockedUser.Password,
		},
	)

	assert.Error(t, err)
}
