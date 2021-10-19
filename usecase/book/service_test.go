package book

import (
	"bookstore.weber.edu/domain"
	"bookstore.weber.edu/domain/book"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tmpUUID = domain.NewID()

// MockBookRepository
type MockBookRepository struct{}

func (mbr *MockBookRepository) ReadAll() ([]book.Book, error) {
	return nil, nil
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
	e := &Service{
		repo: m,
	}
	// act
	a := NewService(m)
	// assert
	assert.IsType(t, e, a, "not typeof Service")
}

func TestNewService_Fail(t *testing.T) {
	// arrange
	m := &MockBookRepository{}
	e := &Service{}
	// act
	a := NewService(m)
	// assert
	assert.NotEqual(t, e.repo, a.repo, "not equal types")
}

func TestService_CreateBook_Success(t *testing.T) {
	// arrange
	m := &MockBookRepository{}
	h := NewService(m)
	eIdType := domain.NewID()
	eId := tmpUUID
	// act
	aId, err := h.CreateBook("123457", "Foobar", "Mr. Foo Bar", 6.66)
	if err != nil {
		t.Fatal("could not create a book in the mock repository")
	}
	// assert
	assert.IsType(t, eIdType, aId, "not type of domain.ID")
	assert.Equal(t, eId, aId)
}
