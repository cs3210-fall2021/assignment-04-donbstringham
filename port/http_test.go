package port

import (
	adapter "bookstore.weber.edu/adapter/book"
	domain "bookstore.weber.edu/domain"
	"bookstore.weber.edu/domain/book"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

type MockBookService struct{}

func NewMockBookService() *MockBookService {
	return &MockBookService{}
}

func (mbs *MockBookService) CreateBook(i string, t string, a string, p float32) (domain.ID, error) {
	return domain.NewID(), nil
}

func (mbs *MockBookService) ListBooks() ([]book.Book, error) {
	bk, _ := book.NewBook("123456", "Foo with Bar", "Mr. Foo Bar", 6.66)
	i := adapter.NewInMemRepository()
	i.Write(bk)
	bks, _ := i.ReadAll()
	return bks, nil
}

func TestCanaryHttp(t *testing.T) {
	assert.True(t, true)
}

func TestNewHttpServer(t *testing.T) {
	// arrange
	m := NewMockBookService()
	// act
	a := NewHttpServer(m)
	// assert
	assert.IsType(t, &HttpServer{}, a)
}

func TestHttpServer_BookHandler(t *testing.T) {
	// arrange
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://localhost:3000/book", nil)
	e := "Foo Bar"
	m := NewMockBookService()
	h := NewHttpServer(m)
	// act
	h.bookGetAll(res, req)
	a, _ := ioutil.ReadAll(res.Body)
	// assert
	assert.Contains(t, string(a), e)
}
