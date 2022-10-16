package models

type Schedule struct {
	Date			string `json:"date"`
	DepartureAt 	string `json:"departure_at"`
	FlightId		int64 `json:"flight_id"`
}

type Flight struct {
	Id       			int64  `json:"-"`
	DepartureId 		int64 `json:"departure_id"`
	ArrivalId 			int64 `json:"arrival_id"`
	AvgDurationMinutes 	int64 `json:"avg_duration_minutes"`

	Departure			Airport `json:"departure"`
	Arrival				Airport `json:"arrival"`

	//Schedules			Schedule `json:"schedules"`
}

type FlightInstance struct {
	Id       			int64  `json:"-"`
	FlightId	 		int64 `json:"flight_id"`
	AircraftId		 	int64 `json:"aircraft_id"`
	Status	 			string `json:"status"`

	Aircraft			Aircraft `json:"aircraft"`
	Flight				Flight `json:"flight"`
}


type FlightInstanceSeat struct {
	Id       			int64  `json:"-"`
	FlightInstanceId	int64 `json:"flight_instance_id"`
	SeatId			 	int64 `json:"seat_id"`
	Cost	 			int64 `json:"cost"`

	Seat				Seat `json:"seat"`
	FlightInstance		FlightInstance `json:"flight_instance"`
}