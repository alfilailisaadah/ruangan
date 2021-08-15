package rents

import (
	"rentRoom/businesses/rents"
	"rentRoom/drivers/databases/users"
	"time"
)

type Rents struct {
	ID          int
	UserId       	int
	JumlahBayar      	int
	TanggalPinjam	string
	User	users.Users
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func fromDomain(domain *rents.Domain) *Rents {
	return &Rents{
		UserId:       domain.UserId,
		JumlahBayar:      domain.JumlahBayar,
		TanggalPinjam: 	domain.TanggalPinjam,
	}
}

func (rec *Rents) ToDomain() rents.Domain {
	return rents.Domain{
		ID:        		rec.ID,
		UserId:       rec.UserId,
		JumlahBayar:      rec.JumlahBayar,
		TanggalPinjam:      rec.TanggalPinjam,
		CreatedAt: 		rec.CreatedAt,
		UpdatedAt: 		rec.UpdatedAt,
	}
}
