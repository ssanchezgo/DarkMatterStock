package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Server es la estructura principal de la API
type Server struct {
	router *chi.Mux
}

// NewServer crea una nueva instancia del servidor
func NewServer() *Server {
	s := &Server{
		router: chi.NewRouter(),
	}
	s.setupRoutes()
	return s
}

// setupRoutes define los endpoints de la API
func (s *Server) setupRoutes() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	// Aquí agregaremos los endpoints más adelante
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("¡Bienvenido a la API DarkMatterStock!"))
	})

	s.router.Get("/stocks", s.handleGetAllStocks)
	s.router.Get("/stocks/{ticker}", s.handleGetStockByTicker)
}

// Run inicia el servidor en el puerto especificado
func (s *Server) Run(port string) error {
	log.Printf("Servidor API iniciado en el puerto %s", port)
	return http.ListenAndServe(":"+port, s.router)
}
