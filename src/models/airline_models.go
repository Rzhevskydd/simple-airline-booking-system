package models

type Airport struct {
	Id			int64 `json:"id"`
	Address 	string `json:"address"`
	Name 		string `json:"name"`
}


type Aircraft struct {
	Id					int64 `json:"id"`
	Name 				string `json:"name"`
	Model 				string `json:"model"`
}

type Seat struct {
	Id						int64 `json:"id"`
	SeatClass 				string `json:"seat_class"`
	AircraftId				int64 `json:"aircraft_id"`
}