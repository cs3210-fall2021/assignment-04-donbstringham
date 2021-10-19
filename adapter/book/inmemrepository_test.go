package book

import (
	"bookstore.weber.edu/domain"
	"bookstore.weber.edu/domain/book"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanaryInMemoryRepository(t *testing.T) {
	assert.True(t, true)
}

func TestNewInMemRepository(t *testing.T) {
	// arrange & act
	harness, err := NewInMemRepository()
	expected, _ := NewInMemRepository()
	// assert
	if err != nil {
		t.Fatal("could not instantiate an InMemRepository", "ERROR")
	}
	assert.IsType(t, expected, harness,"not type of InMemRepository")
}

func TestInMemRepository_Write(t *testing.T) {
	// arrange
	b, _ := book.NewBook("12345", "FooBar", "Foo Bar", 1.11)
	harness, err := NewInMemRepository()
	expected := domain.NewID()
	if err != nil {
		t.Fatal("could not instantiate an InMemRepository", "ERROR")
	}
	// act
	id, err := harness.Write(b)
	if err != nil {
		t.Fatal("could not write a book to an InMemRepository", "ERROR")
	}
	actual, err := harness.ReadAll()
	// assert
	assert.IsType(t, expected, id,"not type of domain.ID")
	found := false
	for _,v := range actual {
		if (v.Isbn == b.Isbn) {
			found = true;
		}
	}
	assert.True(t, found, "book in not in the repository")
}
