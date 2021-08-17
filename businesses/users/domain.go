package users

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Name      string
	Username  string
	Password  string
	UserType  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Login(ctx context.Context, data *Domain) (Domain, string, error)
	CreateToken(ctx context.Context, username, password string) (string, error)
	Store(ctx context.Context, data *Domain) error
}

type Repository interface {
	GetByUsername(ctx context.Context, username string) (Domain, error)
	Store(ctx context.Context, data *Domain) error
}
