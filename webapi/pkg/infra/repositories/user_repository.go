package repositories

import (
	"context"
	"database/sql"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/entities"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type userRepository struct {
	dbConnection *sql.DB
	monitoring   *newrelic.Application
}

func (pst userRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	txn := pst.monitoring.StartTransaction("[UserRepository::FindByEmail]")
	tnxCtx := newrelic.NewContext(context.Background(), txn)

	sql := `SELECT 
								name AS Name, 
								email AS Email, 
								password AS Password
					FROM users
					WHERE email = $1`
	prepare, err := pst.dbConnection.PrepareContext(tnxCtx, sql)
	if err != nil {
		return nil, err
	}

	entity := entities.User{}

	if err := prepare.QueryRowContext(tnxCtx).Scan(
		&entity.Id,
		&entity.Name,
		&entity.Email,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.DeletedAt,
	); err != nil {
		return nil, err
	}
	return &entity, nil
}

func (userRepository) Create(ctx context.Context) (*entities.User, error) {
	return nil, nil
}

func NewUserRepository(dbConnection *sql.DB, monitoring *newrelic.Application) interfaces.IUserRepository {
	return userRepository{
		dbConnection: dbConnection,
		monitoring:   monitoring,
	}
}
