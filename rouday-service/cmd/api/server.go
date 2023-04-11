package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/giofcosta/rouday-api/cmd/api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	Handler		 *Handler
}

// NewServer creates a new server
func NewServer(handler *Handler) *Server {
	return &Server{handler}
}

// RegisterRoutes registers the routes of the server
func (s *Server) RegisterRoutes() {
	prefix := "/v1/api"
	
	docs.SwaggerInfo.Title = "Rouday API"
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Description = "Rouday API Documentation"
    docs.SwaggerInfo.BasePath = prefix
    docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// welcome route
	s.Handler.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to Rouday api!")
    })	

	// ping route
	s.Handler.Router.HandleFunc(prefix + "/ping",  s.Handler.PingHandler).Methods(http.MethodGet)

	// week configuration routes
	s.Handler.Router.HandleFunc(prefix + "/week-configurations/{id}", s.Handler.GetWeekConfiguration).Methods(http.MethodGet)
	s.Handler.Router.HandleFunc(prefix + "/week-configurations", s.Handler.CreateWeekConfiguration).Methods(http.MethodPost)
	s.Handler.Router.HandleFunc(prefix + "/week-configurations/{id}", s.Handler.UpdateWeekConfiguration).Methods(http.MethodPut)
	s.Handler.Router.HandleFunc(prefix + "/week-configurations/{id}/routine/{routine-id}", s.Handler.UpdateWeekConfigurationRoutine).Methods(http.MethodPut)



	// swagger route
	httpSwagger.URL("swagger.json")
	s.Handler.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

// Serve starts the server
func (s *Server) Serve() {
    port := ":8080"
    fmt.Printf("Server is running on %s", port)
    log.Fatal(http.ListenAndServe(port, s.Handler.Router))
}
