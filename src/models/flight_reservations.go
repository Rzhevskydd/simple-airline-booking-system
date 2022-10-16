package models

type Route struct {
	StartAirport 		string `json:"start_airport"`
	FinalAirport 		string `json:"final_airport"`

}

type FlightReservation struct {
	Id       			int64  `json:"-"`
	FlightInstanceId	int64 `json:"flight_instance_id"`
	SeatMap			 	map[int]FlightInstanceSeat `json:"seat_map"`

	FlightInstance		FlightInstance `json:"flight_instance"`
}

type FlightReservationTicketInfo struct {
	Id       				int64  `json:"-"`
	FlightReservationId		int64 `json:"flight_reservation_id"`
	IncludeBaggage			bool `json:"include_baggage"`
	PassengerId				int64 `json:"passenger_id"`

	FlightReservation		FlightReservation `json:"flight_reservation"`
}

