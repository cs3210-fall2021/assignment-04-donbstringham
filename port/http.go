package port

import (
	"bookstore.weber.edu/usecase/book"
	"fmt"
	"log"
	"net/http"
)

type HttpServer struct {
	bookService book.Service
}

func NewHttpServer(b book.Service) *HttpServer {
	return &HttpServer{bookService: b}
}

func (h *HttpServer) Serve() {
	http.HandleFunc("/book", h.bookIndex)
	http.ListenAndServe(":3000", nil)
}

func (h *HttpServer) bookIndex(w http.ResponseWriter, r *http.Request) {
	bks, err := h.bookService.ListBooks()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
