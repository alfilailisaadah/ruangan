package rooms

import (
	"context"
	"rentRoom/businesses"
	"rentRoom/businesses/category"
	"strings"
	"time"
)

type roomsUsecase struct {
	roomsRepository  Repository
	categoryUsecase category.Usecase
	contextTimeout  time.Duration
}

func NewRoomsUsecase(nr Repository, cu category.Usecase, timeout time.Duration) Usecase {
	return &roomsUsecase{
		roomsRepository:  nr,
		categoryUsecase: cu,
		contextTimeout:  timeout,
	}
}

func (nu *roomsUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := nu.roomsRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}
func (nu *roomsUsecase) GetByID(ctx context.Context, roomsId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if roomsId <= 0 {
		return Domain{}, businesses.ErrRoomsIDResource
	}
	res, err := nu.roomsRepository.GetByID(ctx, roomsId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}
func (nu *roomsUsecase) GetByTitle(ctx context.Context, roomsTitle string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if strings.TrimSpace(roomsTitle) == "" {
		return Domain{}, businesses.ErrRoomsTitleResource
	}
	res, err := nu.roomsRepository.GetByTitle(ctx, roomsTitle)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}
func (nu *roomsUsecase) Store(ctx context.Context, ip string, roomsDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	return nil
}
