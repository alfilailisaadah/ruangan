package buildings

import (
	"context"
	"time"
)

type Domain struct {
	ID           	int
	BuildingName    string
	BuildingAddr    string
	CreatedAt   	time.Time
	UpdatedAt    	time.Time
}

type Usecase interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, buildingsId int) (Domain, error)
	GetByTitle(ctx context.Context, buildingsTitle string) (Domain, error)
	Store(ctx context.Context, ip string, buildingsDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, buildingsId int) (Domain, error)
	GetByTitle(ctx context.Context, buildingsTitle string) (Domain, error)
	Store(ctx context.Context, buildingsDomain *Domain) (Domain, error)
	Find(ctx context.Context, rentStatus string) ([]Domain, error)
}
