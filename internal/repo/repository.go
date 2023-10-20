package repo

import (
	"github.com/jmoiron/sqlx"
)

type User interface {
}

type Repository struct {
	User
}

func NewRepo(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
