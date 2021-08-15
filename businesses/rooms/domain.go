package rooms

import (
	"context"
	"time"
)

type Domain struct {
	ID          	int
	RoomName       	string
	RentStatus 		bool
	RentPrice      	int
	BuildingId		int
	BuildingName	string
	CreatedAt   	time.Time
	UpdatedAt   	time.Time
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context,roomsDomain *Domain) (Domain, error)
	GetByActive(ctx context.Context, rentStatus bool) ([]Domain, error)
}

type Repository interface {
	Find(ctx context.Context, rentStatus string) ([]Domain, error)
	Store(ctx context.Context, roomsDomain *Domain) (Domain, error)
	FindByID(id int) (Domain, error)
}

