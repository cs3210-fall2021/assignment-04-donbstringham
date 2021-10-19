package book

import (
	"bookstore.weber.edu/domain"
	"bookstore.weber.edu/domain/book"
)

type Reader interface {
	ReadAll() ([]book.Book, error)
	// ReadByIsbn(isbn string) (*book.Book, error)
}

type Writer interface {
	Write(b *book.Book) (domain.ID, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCases interface {
	ListBooks() ([]book.Book, error)
	CreateBook(i string, t string, a string, p float32) (*domain.ID, error)
}
