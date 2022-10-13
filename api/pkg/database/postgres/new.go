package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func New() (*sql.DB, error) {
	return sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=gocommerce sslmode=disable")
}
