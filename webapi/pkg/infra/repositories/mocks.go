package repositories

import (
	"database/sql"
	"log"
	"os"
	"time"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/entities"
	"webapi/pkg/infra/logger"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type userRepositoryToTest struct {
	repo         interfaces.IUserRepository
	logger       interfaces.ILogger
	dbConnection *sql.DB
	sqlMock      sqlmock.Sqlmock
	newrelic     *newrelic.Application
	mockedUser   entities.User
}

func newUserRepositoryToTest() userRepositoryToTest {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	loggerSpy := logger.NewLoggerSpy()
	newrelic, _ := newrelic.NewApplication(newrelic.ConfigAppName(os.Getenv("APP_NAME")))
	repository := NewUserRepository(loggerSpy, db, newrelic)

	return userRepositoryToTest{
		repo:         repository,
		logger:       loggerSpy,
		dbConnection: db,
		sqlMock:      mock,
		newrelic:     newrelic,
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
