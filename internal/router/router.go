package router

import (
	"github.com/Komilov31/quote-service/internal/handler"
	"github.com/gorilla/mux"
)

func NewRouter(handler *handler.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/quotes", handler.AddQuoteHandler).Methods("POST")
	r.HandleFunc("/quotes", handler.GetQuotesHandler).Methods("GET")
	r.HandleFunc("/quotes/random", handler.GetRandomQuoteHandler).Methods("GET")
	r.HandleFunc("/quotes/{id}", handler.DeleteQuoteHandler).Methods("DELETE")

	return r
}
