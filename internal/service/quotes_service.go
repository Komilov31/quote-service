package service

import "github.com/Komilov31/quote-service/internal/model"

type Service struct {
	repository model.QuoteStore
}

func NewService(repository model.QuoteStore) *Service {
	return &Service{repository: repository}
}

func (s *Service) AddQuote(quote model.Quote) (int, error) {
	return s.repository.AddQuote(quote)
}

func (s *Service) GetAllQuotes() ([]model.Quote, error) {
	return s.repository.GetAllQuotes()
}

func (s *Service) GetRandomQuote() (model.Quote, error) {
	return s.repository.GetRandomQuote()
}

func (s *Service) GetAuthorsQuotes(author string) ([]model.Quote, error) {
	return s.repository.GetAuthorsQuotes(author)
}

func (s *Service) DeleteQuote(id int) error {
	return s.repository.DeleteQuote(id)
}
