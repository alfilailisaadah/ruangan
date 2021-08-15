package rooms

import (
	"context"
	"rentRoom/businesses"
	"time"
)

type roomsUsecase struct {
	roomsRespository Repository
	contextTimeout      time.Duration
}

func NewRoomsUsecase(timeout time.Duration, cr Repository) Usecase {
	return &roomsUsecase{
		contextTimeout:      timeout,
		roomsRespository: cr,
	}
}

func (nu *roomsUsecase) Store(ctx context.Context, newsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	result, err := nu.roomsRespository.Store(ctx, newsDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (cu *roomsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.roomsRespository.Find(ctx, "")
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (cu *roomsUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	if id <= 0 {
		return Domain{}, businesses.ErrIDNotFound
	}

	resp, err := cu.roomsRespository.FindByID(id)
	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}

func (cu *roomsUsecase) GetByActive(ctx context.Context, active bool) ([]Domain, error) {
	findActive := "false"
	if active {
		findActive = "true"
	}
	resp, err := cu.roomsRespository.Find(ctx, findActive)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}
