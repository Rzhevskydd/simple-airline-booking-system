package use_cases

import (
	"booking-system/src/repository"
)

type UseCase struct {
	Flights  FlightsUseCase
}

func NewUseCase(repos *repository.Repositories) *UseCase {
	return &UseCase{
		Flights:  FlightsUseCase{FlightsRepo: repos.Flights},

	}
}