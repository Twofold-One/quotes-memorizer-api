package repository

import (
	"fmt"
	"time"

	quotes_memorizer "github.com/Twofold-One/quotes-memorizer-api"
	"github.com/jmoiron/sqlx"
)

type QuotePostgres struct {
	db *sqlx.DB
}

func NewQuotesPostgres(db *sqlx.DB) *QuotePostgres {
	return &QuotePostgres{db: db}
}

func (r *QuotePostgres) Create(userId int, quote quotes_memorizer.Quote) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createQuoteQuery := fmt.Sprintf(`insert into %s (author, quote, source, sourcetype, created) values ($1, $2, $3, $4, $5) returning id`, quotesTable)
	row := tx.QueryRow(createQuoteQuery, quote.Author, quote.Quote, quote.Source, quote.SourceType, time.Now())
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	
	createUsersQuoteQuery := fmt.Sprintf(`insert into %s (user_id, quote_id) values ($1, $2)`, usersQuotesTable)
	_, err = tx.Exec(createUsersQuoteQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}