package service

import (
	quotes_memorizer "github.com/Twofold-One/quotes-memorizer-api"
	"github.com/Twofold-One/quotes-memorizer-api/pkg/repository"
)

type QuoteService struct {
	repo repository.Quote
}

func NewQuotesService(repo repository.Quote) *QuoteService {
	return &QuoteService{repo: repo}
}

func (s *QuoteService) Create(userId int, quote quotes_memorizer.Quote) (int, error) {
	return s.repo.Create(userId, quote)
}

func (s *QuoteService) GetAll(userID int) ([]quotes_memorizer.Quote, error) {
	return s.repo.GetAll(userID)
}

func (s *QuoteService) GetById(userId, quoteId int) (quotes_memorizer.Quote, error) {
	return s.repo.GetById(userId, quoteId)
}

func (s *QuoteService) Delete(userId, quoteId int) error {
	return s.repo.Delete(userId, quoteId)
}

func (s *QuoteService) Update(userId, quoteId int, input quotes_memorizer.UpdateQuoteInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, quoteId, input)
}
