package rents

import (
	"context"
	"rentRoom/businesses/rooms"
	"strings"
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


func (nu *rentsUsecase) Store(ctx context.Context, rentsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	_, err := nu.rentsRepository.GetById(ctx, rentsDomain.UserId)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{},err
		}
	}

	result, err := nu.rentsRepository.Store(ctx, rentsDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (cu *rentsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.rentsRepository.Find(ctx, "")
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (nu *rentsUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := nu.rentsRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}