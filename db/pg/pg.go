package pg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PgCompany struct {
	Db *sql.DB
}

func InitPgCompany(
	DbUrl string,
) *PgCompany {
	c := &PgCompany{}
	db, err := sql.Open("postgres", DbUrl)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostgreSQL database!")
	c.Db = db
	return c
}

// Execute a query using default context.
func (c *PgCompany) Exec(queryStr string, args ...any) (*sql.Result, error) {
	result, err := c.Db.Exec(queryStr, args...)
	return &result, err
}

// Query a db row using default context.
func (c *PgCompany) QueryRow(queryStr string, args ...any) *sql.Row {
	return c.Db.QueryRow(queryStr, args...)
}
