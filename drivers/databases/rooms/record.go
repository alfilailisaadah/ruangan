package rooms

import (
	roomsUsecase "rentRoom/businesses/rooms"
	"time"
)

type Rooms struct {
	Id        int
	HargaPinjam     int
	StatusPinjam   string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserStat  string
}

func fromDomain(domain *roomsUsecase.Domain) *Rooms {
	return &Rooms{
		Id:       domain.Id,
		HargaPinjam:    domain.HargaPinjam,
		StatusPinjam:  domain.StatusPinjam,
	}
}

func (rec *Rooms) toDomain() roomsUsecase.Domain {
	return roomsUsecase.Domain{
		Id:        rec.Id,
		HargaPinjam:    rec.HargaPinjam,
		StatusPinjam:  rec.StatusPinjam,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
