package rooms

import (
	"context"
	"time"
)

type Domain struct {
	Id         int
	Nama      string
	StatusPinjam    string
	HargaPinjam int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Usecase interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, roomsId int) (Domain, error)
	GetByTitle(ctx context.Context, roomsTitle string) (Domain, error)
	Store(ctx context.Context, ip string, roomsDomain *Domain) error
}

type Repository interface {
	Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetByID(ctx context.Context, roomsId int) (Domain, error)
	GetByTitle(ctx context.Context, roomsTitle string) (Domain, error)
	Store(ctx context.Context, roomsDomain *Domain) error
}
