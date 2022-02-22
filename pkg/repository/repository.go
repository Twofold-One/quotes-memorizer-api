package repository

import (
	quotes_memorizer "github.com/Twofold-One/quotes-memorizer-api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user quotes_memorizer.User) (int, error)
}

type Quotes interface {

}

type Repository struct {
	Authorization
	Quotes
}

// NewSevice is a constructor function.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}