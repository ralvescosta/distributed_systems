package repositories

import (
	"database/sql"
	"log"
	"time"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/entities"
	"webapi/pkg/infra/logger"

	"github.com/DATA-DOG/go-sqlmock"
)

type userRepositoryToTest struct {
	repo         interfaces.IUserRepository
	logger       interfaces.ILogger
	dbConnection *sql.DB
	sqlMock      sqlmock.Sqlmock
	mockedUser   entities.User
}

func newUserRepositoryToTest() userRepositoryToTest {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	loggerSpy := logger.NewLoggerSpy()
	repository := NewUserRepository(loggerSpy, db)

	return userRepositoryToTest{
		repo:         repository,
		logger:       loggerSpy,
		dbConnection: db,
		sqlMock:      mock,
		mockedUser: entities.User{
			Id:        1,
			Name:      "Name",
			Email:     "email@email.com",
			Password:  "password",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
	}
}
