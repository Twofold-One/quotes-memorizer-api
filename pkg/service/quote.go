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