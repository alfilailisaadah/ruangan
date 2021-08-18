package rooms

import (
	"rentRoom/businesses/rooms"
	"rentRoom/drivers/databases/buildings"
	"time"
)

type Rooms struct {
	ID          int
	RoomName       	string
	RentStatus 		bool
	RentPrice      	int
	BuildingId	int
	BuildingName	string
	Building 	buildings.Buildings
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func fromDomain(domain *rooms.Domain) *Rooms {
	return &Rooms{
		ID 		: 		domain.ID,
		RoomName:       domain.RoomName,
		RentStatus: 	domain.RentStatus,
		RentPrice:      domain.RentPrice,
		BuildingId: 	domain.BuildingId,
		BuildingName: 	domain.BuildingName,
	}
}

func (rec *Rooms) ToDomain() rooms.Domain {
	return rooms.Domain{
		ID:        		rec.ID,
		RoomName:       rec.RoomName,
		RentStatus: 	rec.RentStatus,
		RentPrice:      rec.RentPrice,
		BuildingId:		rec.BuildingId,
		BuildingName:	rec.BuildingName,
		CreatedAt: 		rec.CreatedAt,
		UpdatedAt: 		rec.UpdatedAt,
	}
}
