package book

import (
	"bookstore.weber.edu/domain"
	"bookstore.weber.edu/domain/book"
)

type InMemRepository struct {
	repo []book.Book
}

func NewInMemRepository() (*InMemRepository, error) {
	return &InMemRepository{}, nil
}

func (i *InMemRepository) Write(b *book.Book) (domain.ID, error) {
	i.repo = append(i.repo, *b)

	return b.ID, nil
}

func (i *InMemRepository) ReadAll() ([]book.Book, error) {
	return i.repo, nil
}
