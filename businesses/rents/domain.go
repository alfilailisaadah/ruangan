package rents

import (
	"context"
	"time"
)

type Domain struct {
	ID           	int
	UserId   int
	RoomsId    int
	JumlahBayar    int
	TanggalPinjam	string
	CreatedAt   	time.Time
	UpdatedAt    	time.Time
}

type Usecase interface {
	Store(ctx context.Context, ip string, buildingsDomain *Domain) (Domain, error)
}

type Repository interface {
	Store(ctx context.Context, buildingsDomain *Domain) (Domain, error)
}
