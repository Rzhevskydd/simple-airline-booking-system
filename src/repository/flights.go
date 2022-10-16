package repository

import (
	"booking-system/src/models"
	"database/sql"
)

type FlightFilters struct {
	PassengersCount int64
	SeatClass		string
	IncludeBaggage	bool
}

type IFlightRepository interface {
	GetFlightInstances(route models.Route, filters FlightFilters)  []models.FlightInstance
}


type FlightsRepository struct {
	DB *sql.DB
}


func (rep *FlightsRepository) GetFlightInstances(route models.Route, filters FlightFilters)  []models.FlightInstance {


	query := "SELECT " +
				"fi.id, fi.flight_id, fi.aircraft_id, fi.status, " +
				"f.id, f.departure_id, f.arrival_id, f.avg_duration_minutes " +
				"departure.id, departure.address, departure.name " +
				"arrival.id, arrival.address, arrival.name " +
			"FROM flights_instance fi " +
				"JOIN flights f ON fi.flight_id = f.id " +
				"JOIN airports departure ON f.departure_id = departure.id " +
				"JOIN airports arrival ON f.arrival_id = arrival.id " +
				"JOIN flight_instance_seats fis ON fis.flight_instance_id = fi.id " +
				"JOIN seats ON seats.flight_instance_seat_id = fis.id " +
			"WHERE departure.name = $1 AND arrival.name = $2 AND seats.seat_class = $3; "

	rows, _ := rep.DB.Query(query, route.StartAirport, route.FinalAirport, filters.SeatClass)

	var flightInstances []models.FlightInstance
	for rows.Next() {
		flightInstance := models.FlightInstance{}
		flight := models.Flight{}
		fromAirport := models.Airport{}
		toAirport := models.Airport{}

		err := rows.Scan(
			&flightInstance.Id,
			&flightInstance.FlightId,
			&flightInstance.AircraftId,
			&flightInstance.Status,
			&flight.Id,
			&flight.DepartureId,
			&flight.ArrivalId,
			&flight.AvgDurationMinutes,
			&fromAirport.Id,
			&fromAirport.Address,
			&fromAirport.Name,
			&toAirport.Id,
			&toAirport.Address,
			&toAirport.Name,
		)

		flight.Departure = fromAirport
		flight.Arrival = toAirport

		flightInstance.Flight = flight

		if err != nil {
			//return nil, err
		}

		flightInstances = append(flightInstances, flightInstance)
	}

	if err := rows.Close(); err != nil {
		//return nil, err
	}

	return flightInstances
}

