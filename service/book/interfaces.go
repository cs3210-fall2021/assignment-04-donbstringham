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

type ServiceReader interface {
	ListBooks() ([]book.Book, error)
}

type ServiceWriter interface {
	CreateBook(i string, t string, a string, p float32) (domain.ID, error)
}

type Service interface {
	ServiceReader
	ServiceWriter
}
