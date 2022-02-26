package repository

import (
	"fmt"
	"strings"
	"time"

	quotes_memorizer "github.com/Twofold-One/quotes-memorizer-api"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (r *QuotePostgres) GetAll(userId int) ([]quotes_memorizer.Quote, error) {
	var quotes []quotes_memorizer.Quote
	query := fmt.Sprintf(`select q.id, q.author, q.quote, q.source, q.sourcetype, q.created from %s q inner join %s uq on q.id = uq.quote_id where uq.user_id = $1`, quotesTable, usersQuotesTable)
	err := r.db.Select(&quotes, query, userId)

	return quotes, err
}

func (r *QuotePostgres) GetById(userId, quoteId int) (quotes_memorizer.Quote, error) {
	var quote quotes_memorizer.Quote
	query := fmt.Sprintf(`select q.id, q.author, q.quote, q.source, q.sourcetype, q.created from %s q inner join %s uq on q.id = uq.quote_id where uq.user_id = $1 and uq.quote_id = $2`, quotesTable, usersQuotesTable)
	err := r.db.Get(&quote, query, userId, quoteId)

	return quote, err
}

func (r *QuotePostgres) Delete(userId, quoteId int) error {
	query := fmt.Sprintf(`delete from %s q using %s uq where q.id = uq.quote_id and uq.user_id = $1 and uq.quote_id=$2`, quotesTable, usersQuotesTable)
	_, err := r.db.Exec(query, userId, quoteId)

	return err
}

func (r *QuotePostgres) Update(userId, quoteId int, input quotes_memorizer.UpdateQuoteInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Author != nil {
		setValues = append(setValues, fmt.Sprintf(`author=$%d`, argId))
		args = append(args, *input.Author)
		argId++
	}

	if input.Quote != nil {
		setValues = append(setValues, fmt.Sprintf(`quote=$%d`, argId))
		args = append(args, *input.Quote)
		argId++
	}

	if input.Source != nil {
		setValues = append(setValues, fmt.Sprintf(`source=$%d`, argId))
		args = append(args, *input.Source)
		argId++
	}

	if input.SourceType != nil {
		setValues = append(setValues, fmt.Sprintf(`sourcetype=$%d`, argId))
		args = append(args, *input.SourceType)
		argId++
	}

	if input.Created != nil {
		setValues = append(setValues, fmt.Sprintf(`created=$%d`, argId))
		args = append(args, *input.Created)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`update %s q set %s from %s uq where q.id = uq.quote_id and uq.quote_id=$%d and uq.user_id=$%d`, quotesTable, setQuery, usersQuotesTable, argId, argId+1)
	args = append(args, quoteId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}