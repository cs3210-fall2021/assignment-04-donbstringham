package main

import (
	adapter "bookstore.weber.edu/adapter/book"
	http "bookstore.weber.edu/port"
	"bookstore.weber.edu/usecase/book"
)

type BookServices struct {
	Service book.Service
}

func main() {
	bs := *&BookServices{}
	bs.configure()

	webserver := http.NewHttpServer(bs.Service)
	webserver.Serve()
}

func (b *BookServices) configure() {
	a, _ := adapter.NewBookMysqlRepository(
		"donstringham",
		"letmein",
		"143.110.159.170",
		"3306",
		"donstringham",
	)
	s := book.NewService(a)
	b.Service = *s
}
