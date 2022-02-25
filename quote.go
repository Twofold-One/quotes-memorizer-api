package quotes_memorizer

import (
	"errors"
	"time"
)

type Quote struct {
	ID int `json:"id" db:"id"`
	Author string `json:"author" db:"author" binding:"required"`
	Quote string `json:"quote" db:"quote"`
	Source string `json:"source" db:"source"`
	SourceType string `json:"sourcetype" db:"sourcetype"`
	Created time.Time `json:"created" db:"created"`
}

type UsersQuote struct {
	Id int
	UserId int
	QuoteId int
}

type UpdateQuoteInput struct {
	Author *string `json:"author"`
	Quote *string `json:"quote"`
	Source *string `json:"source"`
	SourceType *string `json:"sourcetype"`
	Created *time.Time `json:"created"`
}

func (i UpdateQuoteInput) Validate() error {
	if i.Author == nil && i.Quote == nil && i.Source == nil && i.SourceType == nil && i.Created == nil {
		return errors.New("update structure has no value")
	}
	return nil
}