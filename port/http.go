package port

import (
	service "bookstore.weber.edu/service/book"
	"fmt"
	"log"
	"net/http"
)

type HttpServer struct {
	s service.Service
}

func NewHttpServer(srvc service.Service) *HttpServer {
	return &HttpServer{s: srvc}
}

func (h *HttpServer) BookHandler() {
	http.HandleFunc("/book", h.bookGetAll)
}

func (h *HttpServer) Serve(a string) {
	http.ListenAndServe(a, nil)
}
func (h *HttpServer) bookGetAll(w http.ResponseWriter, r *http.Request) {
	bks, err := h.s.ListBooks()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
