package repository

import (
	quotes_memorizer "github.com/Twofold-One/quotes-memorizer-api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user quotes_memorizer.User) (int, error)
	GetUser(username, password string) (quotes_memorizer.User, error)
}

type Quote interface {
	Create(userId int, quote quotes_memorizer.Quote) (int, error)
}

type Repository struct {
	Authorization
	Quote
}

// NewSevice is a constructor function.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Quote: NewQuotesPostgres(db),
	}
}