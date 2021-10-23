package database

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_Should_Connection_Successfully(t *testing.T) {
	openConnetion = func(driver, connectionString string) (*sql.DB, error) {
		db, _, _ := sqlmock.New()
		return db, nil
	}

	connection, err := GetConnection("some driver", "some connection string")

	assert.NotNil(t, connection)
	assert.NoError(t, err)
}

func Test_Should_Return_Error_When_Try_To_Connection(t *testing.T) {
	openConnetion = func(driver, connectionString string) (*sql.DB, error) {
		return nil, errors.New("error")
	}
	_, err := GetConnection("some driver", "some connection string")

	assert.Error(t, err)
}
