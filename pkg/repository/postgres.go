package repository

import (
	"os"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
	quotesTable = "quotes"
	usersQuotesTable = "users_quotes"
)

// for local env

// type Config struct {
// 	Username string
// 	Password string
// 	Host string
// 	Port string
// 	DBName string
// }

// func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
// 	db, err := sqlx.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

func NewPostgresDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}