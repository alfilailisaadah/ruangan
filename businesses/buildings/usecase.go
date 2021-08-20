package buildings

import (
	"context"
	"rentRoom/businesses"
	"strings"
	"time"
)

type buildingsUsecase struct {
	buildingsRepository  Repository
	// roomsUsecase rooms.Usecase
	contextTimeout  time.Duration
}

func NewBuildingsUsecase(nr Repository, timeout time.Duration) Usecase {
	return &buildingsUsecase{
		buildingsRepository:  nr,
		// roomsUsecase: cu,
		contextTimeout:  timeout,
	}
}

func (nu *buildingsUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := nu.buildingsRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}
func (nu *buildingsUsecase) GetByID(ctx context.Context, buildingsId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if buildingsId <= 0 {
		return Domain{}, businesses.ErrBuildingsIDResource
	}
	res, err := nu.buildingsRepository.GetByID(ctx, buildingsId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}
func (nu *buildingsUsecase) GetByTitle(ctx context.Context, buildingsTitle string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if strings.TrimSpace(buildingsTitle) == "" {
		return Domain{}, businesses.ErrBuildingsTitleResource
	}
	res, err := nu.buildingsRepository.GetByTitle(ctx, buildingsTitle)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}
func (nu *buildingsUsecase) Store(ctx context.Context, buildingsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	result, err := nu.buildingsRepository.Store(ctx, buildingsDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (cu *buildingsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.buildingsRepository.Find(ctx, "")
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}
