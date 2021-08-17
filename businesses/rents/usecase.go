package rents

import (
	"context"
	"rentRoom/businesses/rooms"
	"time"
)

type rentsUsecase struct {
	rentsRepository  Repository
	roomsUsecase rooms.Usecase
	contextTimeout  time.Duration
}

func NewRentsUsecase(nr Repository, cu rooms.Usecase, timeout time.Duration) Usecase {
	return &rentsUsecase{
		rentsRepository:  nr,
		roomsUsecase: cu,
		contextTimeout:  timeout,
	}
}


func (nu *rentsUsecase) Store(ctx context.Context, ip string, rentsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	result, err := nu.rentsRepository.Store(ctx, rentsDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}