package service

import (
	quotes_memorizer "github.com/Twofold-One/quotes-memorizer-api"
	"github.com/Twofold-One/quotes-memorizer-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user quotes_memorizer.User) (int, error)
}

type Quotes interface {

}

type Service struct {
	Authorization
	Quotes
}

// NewSevice is a constructor function.
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}