package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Komilov31/quote-service/internal/handler"
	"github.com/Komilov31/quote-service/internal/logger"
	"github.com/Komilov31/quote-service/internal/model"
	"github.com/Komilov31/quote-service/internal/repository"
	"github.com/Komilov31/quote-service/internal/router"
	"github.com/Komilov31/quote-service/internal/service"
)

var (
	repo = repository.NewRepository()
	log  = logger.NewLogger()
	svc  = service.NewService(repo)
	hdlr = handler.NewHandler(svc, log)
	r    = router.NewRouter(hdlr)
)

func TestAddQuoteHandler(t *testing.T) {

	quotes := []model.Quote{
		{Author: "Confucius", Text: "Life is simple, but we insist on making it complicated."},
		{Author: "Шекспир", Text: "Быть или не быть, вот в чем вопрос"},
		{Author: "Стетхем", Text: "Лучше быть, чем не быть"},
	}
	for _, quote := range quotes {
		body, _ := json.Marshal(quote)

		req := httptest.NewRequest("POST", "/quotes", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
		}
	}
}

func TestGetAllQuotesHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/quotes", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	quotes, _ := svc.GetAllQuotes()
	if len(quotes) != 3 {
		t.Errorf("expected 3 quote, got %d", len(quotes))
	}
}

func TestGetQuotesByAuthor(t *testing.T) {
	req := httptest.NewRequest("GET", "/quotes?author=Confucius", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var quotes []model.Quote
	json.Unmarshal(w.Body.Bytes(), &quotes)
	for _, quote := range quotes {
		if quote.Author != "Confucius" {
			t.Errorf("Expected Confucius as author got %s", quote.Author)
		}
	}
}

func TestGetRandomQuote(t *testing.T) {
	req := httptest.NewRequest("GET", "/quotes/random", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var quote model.Quote
	json.NewDecoder(w.Body).Decode(&quote)
	if quote.Author == "" {
		t.Errorf("Test of getting random quote failed")
	}
}

func TestDeleteQuote(t *testing.T) {

	req := httptest.NewRequest("DELETE", "/quotes/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	quotes, _ := svc.GetAllQuotes()
	if len(quotes) != 2 {
		t.Errorf("expected length 2 after deleting, but got %d", len(quotes))
	}

	_, err := repo.GetQuoteByID(1)
	if err == nil {
		t.Errorf("quote had to be deleted, but did not")
	}
}
