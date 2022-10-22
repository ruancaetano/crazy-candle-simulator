package internal

import (
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal/controllers"
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal/entities"
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal/repositories"
	"net/http"
)

type Server struct {
	newCandleChan chan entities.Candle
	repository    *repositories.MongoRepository
}

func NewServer(channel chan entities.Candle, repository *repositories.MongoRepository) *Server {
	server := &Server{
		newCandleChan: channel,
		repository:    repository,
	}

	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	websocketController := controllers.NewWebSocketController(s.newCandleChan)
	getCandlesRepository := controllers.NewGetCandlesController(s.repository)

	http.HandleFunc("/ws", websocketController.Execute)
	http.HandleFunc("/candles", getCandlesRepository.Execute)
}

func (*Server) Listen(address string) error {
	return http.ListenAndServe(address, nil)
}
