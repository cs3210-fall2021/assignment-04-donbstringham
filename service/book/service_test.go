package book

import (
	"bookstore.weber.edu/domain"
	"bookstore.weber.edu/domain/book"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tmpUUID = domain.NewID()

// MockBookRepository
type MockBookRepository struct {
	r []book.Book
}

func (mbr *MockBookRepository) ReadAll() ([]book.Book, error) {
	b := book.Book{
		Isbn:   "123456",
		Title:  "Foo Bar",
		Author: "John Doe",
		Price:  1.23,
	}
	mbr.r = append(mbr.r, b)
	return mbr.r, nil
}

func (mbr *MockBookRepository) Write(b *book.Book) (domain.ID, error) {
	return tmpUUID, nil
}

// Unit Tests
func TestCanaryBookService(t *testing.T) {
	assert.True(t, true)
}

func TestNewService_Success(t *testing.T) {
	// arrange
	m := &MockBookRepository{}
	e := &BookService{
		repo: m,
	}
	// act
	a := NewBookService(m)
	// assert
	assert.IsType(t, e, a, "not typeof Service")
}

func TestNewService_Fail(t *testing.T) {
	// arrange
	m := &MockBookRepository{}
	e := &BookService{}
	// act
	a := NewBookService(m)
	// assert
	assert.NotEqual(t, e.repo, a.repo, "not equal types")
}

func TestService_CreateBook_Success(t *testing.T) {
	// arrange
	m := &MockBookRepository{}
	h := NewBookService(m)
	eIdType := domain.NewID()
	eId := tmpUUID
	// act
	a, err := h.ListBooks()
	// assert
	assert.Equal(t, nil, err)
	assert.IsType(t, eIdType, a[0].ID, "not type of domain.ID")
	assert.Equal(t, eId, a[0].ID)
}
