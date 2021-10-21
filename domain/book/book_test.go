package book

import (
	"bookstore.weber.edu/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBook_Success(t *testing.T) {
	// arrange
	e := &Book{
		Isbn:   "123456",
		Title:  "Test Book",
		Author: "Foo Bar",
		Price:  11.99,
	}
	// act
	a, _ := NewBook("123456", "Test Book", "Foo Bar", 11.99)
	// assert
	assert.NotEqual(t, e.ID, a.ID)
	assert.Equal(t, e.Isbn, a.Isbn)
	assert.Equal(t, e.Title, a.Title)
	assert.Equal(t, e.Author, a.Author)
	assert.Equal(t, e.Price, a.Price)
}
func TestNewBook_Fail(t *testing.T) {
	// arrange
	e := &Book{
		Isbn:   "123456",
		Title:  "Test Book",
		Author: "Foo Bar",
		Price:  11.99,
	}
	// act
	a, _ := NewBook("12346", "Test Book", "Fo Bar", 11.98)
	// assert
	assert.NotEqual(t, e.ID, a.ID)
	assert.NotEqual(t, e.Isbn, a.Isbn)
	assert.Equal(t, e.Title, a.Title)
	assert.NotEqual(t, e.Author, a.Author)
	assert.NotEqual(t, e.Price, a.Price)
}
func TestBook_Validate_Success(t *testing.T) {
	// arrange
	h := &Book{
		Isbn:   "",
		Title:  "Test Book",
		Author: "Foo Bar",
		Price:  11.99,
	}
	// act
	a := h.Validate()
	// assert
	assert.Equal(t, domain.ErrInvalidEntity, a)
}
