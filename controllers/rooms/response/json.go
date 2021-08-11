package response

import (
	"rentRoom/businesses/rooms"
	"time"
)

type Rooms struct {
	Id        int       `json:"id"`
	StatusPinjam     string    `json:"statusPinjam"`
	HargaPinjam   int    `json:"hargaPinjam"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain rooms.Domain) Rooms {
	return Rooms{
		Id:        domain.Id,
		StatusPinjam:     domain.StatusPinjam,
		HargaPinjam:   domain.HargaPinjam,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
