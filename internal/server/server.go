package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/stellafff25/Lab5/db/sqlc"
	"github.com/stellafff25/Lab5/internal/server/handlers"
)

type Server struct {
	router *mux.Router
	store  db.Store
}

func NewServer(store db.Store) *Server {
	s := &Server{
		router: mux.NewRouter(),
		store:  store,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.HandleFunc("/health", s.handleHealth).Methods("GET")

	// Create order handler
	orderHandler := handlers.NewOrderHandler(s.store)

	// Order routes
	s.router.HandleFunc("/orders", orderHandler.GetAllOrders).Methods("GET")
	s.router.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	s.router.HandleFunc("/orders/{id}", orderHandler.GetOrder).Methods("GET")
	s.router.HandleFunc("/orders/{id}", orderHandler.UpdateOrder).Methods("PUT")
	s.router.HandleFunc("/orders/{id}", orderHandler.DeleteOrder).Methods("DELETE")
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (s *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.router))
}
