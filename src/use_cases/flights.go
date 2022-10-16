package use_cases

import (
	"booking-system/src/models"
	"booking-system/src/repository"
	"net/url"
	"strconv"
)


type IFlightsUseCase interface {
	GetFlightInstances(route models.Route, filters *repository.FlightFilters)  ([]models.FlightInstance, error)

}

type FlightsUseCase struct {
	FlightsRepo repository.FlightsRepository
}


func (u *FlightsUseCase) GetFlightInstances(
	params url.Values,
) ([]models.FlightInstance, error) {

	fromAirport := params.Get("from")
	toAirport := params.Get("to")

	passengersCount := params.Get("passengersCount")
	intPassengersCount, _ := strconv.Atoi(passengersCount)

	seatClass := params.Get("seatClass")
	includeBaggage := params.Get("includeBaggage")

	route := models.Route{
		StartAirport: fromAirport,
		FinalAirport: toAirport,
	}

	filters := repository.FlightFilters{
		PassengersCount: int64(intPassengersCount),
		SeatClass:       seatClass,
		IncludeBaggage:  stringToBool(includeBaggage),
	}

	flightsInstances := u.FlightsRepo.GetFlightInstances(route, filters)

	return flightsInstances, nil
}