package service

import (
	quotes_memorizer "github.com/Twofold-One/quotes-memorizer-api"
	"github.com/Twofold-One/quotes-memorizer-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user quotes_memorizer.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Quote interface {
	Create(userId int, quote quotes_memorizer.Quote) (int, error)
}

type Service struct {
	Authorization
	Quote
}

// NewSevice is a constructor function.
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Quote: NewQuotesService(repos.Quote),
	}
}