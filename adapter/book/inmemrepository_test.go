package book

import (
	"bookstore.weber.edu/domain"
	"bookstore.weber.edu/domain/book"
	"reflect"
	"testing"
)

func TestNewInMemRepository(t *testing.T) {
	// arrange & act
	harness, err := NewInMemRepository()
	expected, _ := NewInMemRepository()
	// assert
	if err != nil {
		t.Fatal("could not instantiate an InMemRepository", "ERROR")
	}
	if reflect.TypeOf(expected) != reflect.TypeOf(harness) {
		t.Fatal("not type of InMemRepository", "ERROR")
	}
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
	// assert
	if err != nil {
		t.Fatal("could not write a book to an InMemRepository", "ERROR")
	}
	if reflect.TypeOf(expected) != reflect.TypeOf(id) {
		t.Fatal("not type of domain.ID", "ERROR")
	}
}
