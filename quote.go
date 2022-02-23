package quotes_memorizer

import "time"

type Quote struct {
	ID int `json:"id"`
	Author string `json:"author" binding:"required"`
	Quote string `json:"quote"`
	Source string `json:"source"`
	SourceType string `json:"sourcetype"`
	Created time.Time `json:"created"`
}

type UsersQuote struct {
	Id int
	UserId int
	QuoteId int
}
