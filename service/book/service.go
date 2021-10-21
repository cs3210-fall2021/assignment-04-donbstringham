package book

import (
	"bookstore.weber.edu/domain/book"

	domain "bookstore.weber.edu/domain"
)

type BookService struct {
	repo Repository
}

func NewBookService(r Repository) *BookService {
	return &BookService{repo: r}
}

func (s *BookService) CreateBook(i string, t string, a string, p float32) (domain.ID, error) {
	b, err := book.NewBook(i, t, a, p)
	if err != nil {
		return b.ID, err
	}
	id, err := s.repo.Write(b)
	if err != nil {
		return b.ID, err
	}
	return id, nil
}

func (s *BookService) ListBooks() ([]book.Book, error) {
	bks, err := s.repo.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(bks) == 0 {
		return nil, domain.ErrNotFound
	}
	return bks, nil
}
