package repository

import "database/sql"

type Repositories struct {
	Flights  FlightsRepository
}

func CreateRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Flights:  FlightsRepository{DB: db},

	}
}
