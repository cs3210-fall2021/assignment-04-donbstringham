package main

import (
	adapter "bookstore.weber.edu/adapter/book"
	"bookstore.weber.edu/port"
	"bookstore.weber.edu/service/book"
)

func main() {
	a := confAdapter()
	s := confService(a)
	w := port.NewHttpServer(s)
	w.BookHandler()
	w.Serve(":3000")
}

func confAdapter() *adapter.MysqlRepository {
	a, _ := adapter.NewBookMysqlRepository(
		"donstringham",
		"letmein",
		"143.110.159.170",
		"3306",
		"donstringham",
	)
	return a
}

func confService(a *adapter.MysqlRepository) *book.BookService {
	s := book.NewBookService(a)
	return s
}
