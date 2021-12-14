package repositories

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/entities"
	"webapi/pkg/infra/logger"
	"webapi/pkg/infra/telemetry"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type userRepositoryToTest struct {
	repo         interfaces.IUserRepository
	logger       interfaces.ILogger
	dbConnection *sql.DB
	sqlMock      sqlmock.Sqlmock
	mockedUser   entities.User
	telemetry    telemetry.ITelemetry
}

func newUserRepositoryToTest() userRepositoryToTest {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	loggerSpy := logger.NewLoggerSpy()
	telemetry := newTelemetrySpy()
	repository := NewUserRepository(loggerSpy, db, telemetry)

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
		telemetry: telemetry,
	}
}

type telemetrySpy struct{}

func (telemetrySpy) GinMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func (telemetrySpy) InstrumentQuery(ctx context.Context, sqlType string, sql string) opentracing.Span {
	return opentracing.StartSpan("")
}
func (telemetrySpy) InstrumentGRPCClient(ctx context.Context, clientName string) (opentracing.Span, context.Context) {
	return nil, nil
}
func (telemetrySpy) InstrumentAMQPPublisher(ctx context.Context, exchangeName, queueName string) (opentracing.Span, context.Context) {
	return nil, nil
}
func (telemetrySpy) StartSpanFromRequest(header http.Header) opentracing.Span {
	return opentracing.StartSpan("")
}
func (telemetrySpy) Inject(span opentracing.Span, request *http.Request) error {
	return nil
}
func (telemetrySpy) InjectAMQPHeader(header map[string]interface{}, ctx context.Context) error {
	return nil
}
func (telemetrySpy) GetTraceparenteFromSpan(span opentracing.Span) string {
	return ""
}
func (telemetrySpy) Extract(header http.Header) (opentracing.SpanContext, error) {
	return nil, nil
}
func (telemetrySpy) Dispatch() {}
func (telemetrySpy) GetTracer() opentracing.Tracer {
	return nil
}

func newTelemetrySpy() telemetry.ITelemetry {
	return telemetrySpy{}
}
