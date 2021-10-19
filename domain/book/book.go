package book

import "bookstore.weber.edu/domain"

// Book ...
type Book struct {
	ID     domain.ID `json:"ID"`
	Isbn   string    `json:"isbn"`
	Title  string    `json:"Title" validate:"required"`
	Author string    `json:"Author" validate:"required"`
	Price  float32   `json:"Price"`
}

func NewBook(i string, t string, a string, p float32) (*Book, error) {
	b := &Book{
		ID:     domain.NewID(),
		Isbn:   i,
		Title:  t,
		Author: a,
		Price:  p,
	}

	err := b.Validate()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Book) Validate() error {
	if b.Isbn == "" || b.Title == "" || b.Author == "" || b.Price <= 0 {
		return domain.ErrInvalidEntity
	}
	return nil
}
