package rents

import (
	"rentRoom/businesses/rents"
	"rentRoom/drivers/databases/rooms"
	"rentRoom/drivers/databases/users"
	"time"
)

type Rents struct {
	ID          int
	UserId       	int
	RoomId       	int
	JumlahBayar      	int
	TanggalPinjam	string
	StatusPinjam	bool
	User	users.Users
	Room	rooms.Rooms
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func fromDomain(domain *rents.Domain) *Rents {
	return &Rents{
		UserId:       domain.UserId,
		RoomId:       domain.RoomId,
		JumlahBayar:      domain.JumlahBayar,
		TanggalPinjam: 	domain.TanggalPinjam,
		StatusPinjam: domain.StatusPinjam,
	}
}

func (rec *Rents) ToDomain() rents.Domain {
	return rents.Domain{
		ID:        		rec.ID,
		UserId:       rec.UserId,
		RoomId:       rec.RoomId,
		JumlahBayar:      rec.JumlahBayar,
		TanggalPinjam:      rec.TanggalPinjam,
		StatusPinjam: rec.StatusPinjam,
		CreatedAt: 		rec.CreatedAt,
		UpdatedAt: 		rec.UpdatedAt,
	}
}
