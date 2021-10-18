package repositories

import (
	"context"
	"database/sql"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/entities"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type userRepository struct {
	logger       interfaces.ILogger
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
		pst.logger.Error(err.Error())
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
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		pst.logger.Error(err.Error())
		return nil, err
	}
	return &entity, nil
}

func (pst userRepository) Create(ctx context.Context, dto dtos.CreateUserDto) (*entities.User, error) {
	sql := `INSERT INTO users
								(name, email, password) 
					VALUES
								($1, $2, $3) 
					RETURNING *`

	prepare, err := pst.dbConnection.PrepareContext(ctx, sql)
	if err != nil {
		pst.logger.Error(err.Error())
		return nil, err
	}

	entity := entities.User{}

	row := prepare.QueryRowContext(ctx, dto.Name, dto.Email, dto.Password)
	if row == nil {
		return nil, nil
	}

	if err = row.Scan(
		&entity.Id,
		&entity.Name,
		&entity.Email,
		&entity.Password,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.DeletedAt,
	); err != nil {
		pst.logger.Error(err.Error())
		return nil, err
	}

	return &entity, nil
}

func NewUserRepository(logger interfaces.ILogger, dbConnection *sql.DB, monitoring *newrelic.Application) interfaces.IUserRepository {
	return userRepository{
		logger:       logger,
		dbConnection: dbConnection,
		monitoring:   monitoring,
	}
}
