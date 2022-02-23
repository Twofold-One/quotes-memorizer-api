package repository

import (
	"fmt"

	quotes_memorizer "github.com/Twofold-One/quotes-memorizer-api"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user quotes_memorizer.User) (int, error) {
	var id int
	query := fmt.Sprintf(`insert into %s (name, username, password_hash) values ($1, $2, $3) returning id`, usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (quotes_memorizer.User, error) {
	var user quotes_memorizer.User

	query := fmt.Sprintf(`select id from %s where username=$1 and password_hash=$2`, usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}