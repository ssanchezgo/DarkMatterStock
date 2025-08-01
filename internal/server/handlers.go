package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// handleGetAllStocks es el manejador para el endpoint GET /stocks
func (s *Server) handleGetAllStocks(w http.ResponseWriter, r *http.Request) {
	// Lógica para obtener todos los stocks de la base de datos
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Este endpoint devolverá todos los stocks"))
}

// handleGetStockByTicker es el manejador para el endpoint GET /stocks/{ticker}
func (s *Server) handleGetStockByTicker(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")
	// Lógica para obtener un stock por su ticker
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Este endpoint devolverá el stock para el ticker: " + ticker))
}

// respondWithJSON es una función auxiliar para enviar respuestas en formato JSON
func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}
