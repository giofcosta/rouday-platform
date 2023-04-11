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

// @Summary Get Week Configuration
// @Description Gets a week configuration by id
// @Tags week-configurations
// @Accept  json
// @Produce  json
// @Param id path string true "Week Configuration ID"
// @Success 200 {object} models.WeekConfiguration
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /week-configurations/{id} [get]
func (h *Handler) GetWeekConfiguration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	config, err := h.App.WeekConfigurationRepository.GetByID(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if config == nil {
		log.Println("config not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(config)
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

	err = h.App.WeekConfigurationRepository.Create(&config)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// @Summary Update Week Configuration
// @Description Updates a week configuration
// @Tags week-configurations
// @Accept  json
// @Produce  json
// @Param id path string true "Week Configuration ID"
// @Param weekConfiguration body models.WeekConfiguration true "Update Week Configuration"
// @Success 200 {object} models.WeekConfiguration
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /week-configurations/{id} [put]
func (h *Handler) UpdateWeekConfiguration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var config models.WeekConfiguration
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		log.Println(err)	
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	config.ID = id

	err = h.App.WeekConfigurationRepository.Update(&config)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateWeekConfigurationRoutine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	routineID := vars["routine_id"]

	var routine models.Routine
	err := json.NewDecoder(r.Body).Decode(&routine)
	if err != nil {
		log.Println(err)	
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	routine.ID = routineID

	err = h.App.WeekConfigurationRepository.UpdateRoutine(id, &routine)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

