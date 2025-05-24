package api

import (
	"log"
	"net/http"

	"github.com/Komilov31/quote-service/internal/handler"
	"github.com/Komilov31/quote-service/internal/logger"
	"github.com/Komilov31/quote-service/internal/repository"
	"github.com/Komilov31/quote-service/internal/router"
	"github.com/Komilov31/quote-service/internal/service"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Run() error {

	logger := logger.NewLogger()
	repository := repository.NewRepository()
	service := service.NewService(repository)
	handler := handler.NewHandler(service, logger)

	router := router.NewRouter(handler)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
