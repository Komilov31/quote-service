package repository

import (
	"errors"
	"math/rand"
	"sync"

	"github.com/Komilov31/quote-service/internal/model"
)

type Repository struct {
	store      map[int]model.Quote
	lastItemID int
	mu         sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{
		store:      make(map[int]model.Quote),
		lastItemID: 0,
		mu:         sync.RWMutex{},
	}
}

func (r *Repository) AddQuote(quote model.Quote) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastItemID++
	quote.ID = r.lastItemID
	r.store[r.lastItemID] = quote

	return r.lastItemID, nil
}

func (r *Repository) GetAllQuotes() ([]model.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	quotes := make([]model.Quote, 0)
	for _, quote := range r.store {
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

func (r *Repository) GetRandomQuote() (model.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.store) == 0 {
		return model.Quote{}, errors.New("no quotes in storage")
	}

	randomNumber := rand.Int63n(int64(r.lastItemID)) + 1
	randomQuote, ok := r.store[int(randomNumber)]
	for !ok {
		randomNumber = rand.Int63n(int64(r.lastItemID)) + 1
		randomQuote, ok = r.store[int(randomNumber)]
	}

	return randomQuote, nil
}

func (r *Repository) GetAuthorsQuotes(author string) ([]model.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	authorsQuotes := make([]model.Quote, 0)
	for _, quote := range r.store {
		if quote.Author == author {
			authorsQuotes = append(authorsQuotes, quote)
		}
	}

	return authorsQuotes, nil
}

func (r *Repository) DeleteQuote(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.store[id]
	if !ok {
		return errors.New("no quote with such id")
	}

	delete(r.store, id)
	return nil
}

func (r *Repository) GetQuoteByID(id int) (model.Quote, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	quote, ok := r.store[id]
	if !ok {
		return model.Quote{}, errors.New("no quote with such id")
	}

	return quote, nil
}
