package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "library"
	password = "root"
)

func CreateDB() (*sql.DB, error) {
	conn := fmt.Sprintf(`host=%s port=%d user=%s dbname=%s password=%s 
	sslmode=disable`, host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	return db, err
}
