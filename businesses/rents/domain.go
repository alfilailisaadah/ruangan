package rents

import (
	"context"
	"time"
)

type Domain struct {
	ID           	int
	UserId   int
	RoomId    int
	JumlahBayar    int
	TanggalPinjam	string
	CreatedAt   	time.Time
	UpdatedAt    	time.Time
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	Store(ctx context.Context, rentsDomain *Domain) (Domain, error)
}

type Repository interface {
	Find(ctx context.Context, rentStatus string) ([]Domain, error)
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetById(ctx context.Context, userId int) (Domain, error)
	GetRoomById(ctx context.Context, roomId int) (Domain, error)
	Store(ctx context.Context, rentsDomain *Domain) (Domain, error)
}
