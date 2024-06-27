// This file contains the repository implementation layer.
package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Repository struct {
	Db *sql.DB
}

type NewRepositoryOptions struct {
	Dsn string
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	db, err := sql.Open("postgres", opts.Dsn)
	if err != nil {
		log.Panic("can't connect db postgres")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panic("can't ping db postgres")
		panic(err)
	}

	return &Repository{
		Db: db,
	}
}
