package service

import "github.com/Twofold-One/quotes-memorizer-api/pkg/repository"

type Authorization interface {

}

type Quotes interface {

}

type Service struct {
	Authorization
	Quotes
}

// NewSevice is a constructor function.
func NewService(repos *repository.Repository) *Service {
	return &Service{}
}