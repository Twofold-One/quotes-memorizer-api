package repository

type Authorization interface {

}

type Quotes interface {

}

type Repository struct {
	Authorization
	Quotes
}

// NewSevice is a constructor function.
func NewRepository() *Repository {
	return &Repository{}
}