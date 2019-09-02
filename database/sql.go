package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const (
	SqlTypePostgres = "postgres"
)


func OpenSql(connectionType, connectionString string) *sql.DB {
	db, err := sql.Open(connectionType, connectionString)
	if err != nil {
		log.Fatalf("ERROR - Opening sql connection. %v\n", err)
	}

	return db
}
