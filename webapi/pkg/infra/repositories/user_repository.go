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
	// txn := pst.monitoring.StartTransaction("postgresQuery")
	// defer txn.End()
	// tnxCtx := newrelic.NewContext(ctx, txn)

	sql := `SELECT 
								id as Id,
								name AS Name, 
								email AS Email, 
								password AS Password,
								created_at AS CreatedAt,
								updated_at AS UpdatedAt,
								deleted_at AS DeletedAt
					FROM users
					WHERE email = $1
					AND deleted_at IS NULL`

	prepare, err := pst.dbConnection.PrepareContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	entity := entities.User{}

	row := prepare.QueryRowContext(ctx, email)
	if row == nil {
		return nil, nil
	}

	if err := row.Scan(
		&entity.Id,
		&entity.Name,
		&entity.Email,
		&entity.Password,
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
