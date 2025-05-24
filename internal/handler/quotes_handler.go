package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Komilov31/quote-service/internal/logger"
	"github.com/Komilov31/quote-service/internal/model"
	"github.com/Komilov31/quote-service/internal/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	quoteService model.QuoteService
	logger       logger.Logger
}

func NewHandler(quoteService model.QuoteService, logger logger.Logger) *Handler {
	return &Handler{
		quoteService: quoteService,
		logger:       logger,
	}
}

func (h *Handler) AddQuoteHandler(w http.ResponseWriter, r *http.Request) {
	var quote model.Quote
	if err := utils.ParseJson(r, &quote); err != nil {
		h.logger.ErrorLogger.Println("invalid POST request to", r.URL.Path)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	id, err := h.quoteService.AddQuote(quote)
	if err != nil {
		h.logger.ErrorLogger.Println("server error to POST request to", r.URL.Path)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]int{"id": id})
	h.logger.InfoLogger.Println("processed POST request and added to storage")
}

func (h *Handler) GetQuotesHandler(w http.ResponseWriter, r *http.Request) {
	var quotes []model.Quote
	var err error

	author := r.URL.Query().Get("author")
	if author == "" {
		quotes, err = h.quoteService.GetAllQuotes()
	} else {
		quotes, err = h.quoteService.GetAuthorsQuotes(author)
	}
	if err != nil {
		h.logger.ErrorLogger.Println("internal server error to GET request to", r.URL.Path)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, quotes)
	h.logger.InfoLogger.Println("processed GET request quotes were returned")
}

func (h *Handler) GetRandomQuoteHandler(w http.ResponseWriter, r *http.Request) {
	randomQuote, err := h.quoteService.GetRandomQuote()
	if err != nil {
		h.logger.ErrorLogger.Println("internal server error to", r.URL.Path)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, randomQuote)
	h.logger.InfoLogger.Println("processed GET request and random quote was returned")
}

func (h *Handler) DeleteQuoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		h.logger.ErrorLogger.Println("invalid id was given to", r.URL.Path)
		utils.WriteError(w, http.StatusInternalServerError, errors.New("invalid id"))
		return
	}

	err = h.quoteService.DeleteQuote(id)
	if err != nil {
		h.logger.ErrorLogger.Println("no user with id provided to path", r.URL.Path)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]int{"id": id})
	h.logger.InfoLogger.Println("processed DELETE request and deleted quote")
}
