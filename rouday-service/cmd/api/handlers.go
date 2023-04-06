package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/giofcosta/rouday-api/cmd/api/data/models"
	"github.com/gorilla/mux"
)

type Handler struct {
	App *Application
	Router *mux.Router
}

func NewHandler(app *Application, router *mux.Router) *Handler {
	return &Handler{app, router}
}


// @Summary Ping
// @Description Simple ping method for testing purposes
// @Tags ping
// @Produce  json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "pong"}`))
}

// @Summary Create Week Configuration
// @Description Creates a new week configuration
// @Tags week-configurations
// @Accept  json
// @Produce  json
// @Param weekConfiguration body models.WeekConfiguration true "Create Week Configuration"
// @Success 201 {object} models.WeekConfiguration
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /week-configurations [post]
func (h *Handler) CreateWeekConfiguration(w http.ResponseWriter, r *http.Request) {
	var config models.WeekConfiguration
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		log.Println(err)	
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.WeekConfigurationRepository.CreateWeekConfiguration(&config)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

