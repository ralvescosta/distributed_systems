package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
)

func GetConnection(host string, port string, user, password, dbName string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, err := sql.Open("nrpostgres", connectionString)
	if err != nil {
		log.Printf("error while connect to database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("error while check database connection: %v", err)
		return nil, err
	}

	return db, nil
}
