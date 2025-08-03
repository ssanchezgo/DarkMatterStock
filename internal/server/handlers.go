package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ssanchezgo/DarkMatterStock/internal/db"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

// handleGetAllStocks es el manejador para el endpoint GET /stocks
func (s *Server) handleGetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := db.GetAllStocks(r.Context())
	if err != nil {
		log.Printf("Error al obtener todos los stocks: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, stocks)
}

// handleGetStockByTicker es el manejador para el endpoint GET /stocks/{ticker}
func (s *Server) handleGetStockByTicker(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")
	if ticker == "" {
		http.Error(w, "Debe especificar un ticker", http.StatusBadRequest)
		return
	}

	stock, err := db.GetStockByTicker(r.Context(), ticker)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			log.Printf("Error al obtener el stock por ticker %s: %v", ticker, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	respondWithJSON(w, http.StatusOK, stock)
}

// respondWithJSON es una función auxiliar para enviar respuestas en formato JSON
func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}
