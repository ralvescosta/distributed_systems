package database

import (
	"database/sql"
	"log"

	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
)

var openConnetion = sql.Open

func GetConnection(driver, connectionString string) (*sql.DB, error) {
	db, err := openConnetion(driver, connectionString)
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
