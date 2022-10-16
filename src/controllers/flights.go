package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	u "booking-system/src/use_cases"
)


type ApiFlightsHandler struct {
	Flights u.FlightsUseCase
}

func HandleFlightsRoutes(m *mux.Router, usecase *u.UseCase) {
	api := ApiFlightsHandler{Flights: usecase.Flights}

	m.HandleFunc("/", api.HandleGetFlights ).Methods(http.MethodGet)
}

func (api *ApiFlightsHandler) HandleGetFlights(w http.ResponseWriter, r *http.Request) {
	defer CloseBody(w, r)
	w.Header().Set("Content-Type", "application/json")

	flightsInstances, err := api.Flights.GetFlightInstances(r.URL.Query())

	if err != nil {
		NewError(w, http.StatusBadRequest, "Unable to get flights")
		return
	}


	ResponseJson(w, http.StatusOK, flightsInstances)
}
