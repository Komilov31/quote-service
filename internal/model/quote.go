package model

type Quote struct {
	ID     int    `json:"id"`
	Text   string `json:"quote"`
	Author string `json:"author"`
}

type QuoteStore interface {
	AddQuote(quote Quote) (int, error)
	GetAllQuotes() ([]Quote, error)
	GetRandomQuote() (Quote, error)
	GetAuthorsQuotes(string) ([]Quote, error)
	DeleteQuote(id int) error
}

type QuoteService interface {
	AddQuote(quote Quote) (int, error)
	GetAllQuotes() ([]Quote, error)
	GetRandomQuote() (Quote, error)
	GetAuthorsQuotes(string) ([]Quote, error)
	DeleteQuote(id int) error
}
