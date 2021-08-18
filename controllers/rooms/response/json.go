package response

import (
	"rentRoom/businesses/rooms"
	"time"
)

type Rooms struct {
	ID         int       `json:"id"`
	RoomName    string	  `json:"room_name"`
	RentStatus 	bool	  `json:"rent_status"`
	RentPrice   int       `json:"rent_price"`
	BuildingId   	int		`json:"building_id"`
	BuildingName   	string		`json:"building_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain rooms.Domain) Rooms {
	return Rooms{
		ID:          	domain.ID,
		RoomName:       domain.RoomName,
		RentStatus: 	domain.RentStatus,
		RentPrice:      domain.RentPrice,
		BuildingId: 	domain.BuildingId,
		BuildingName: 	domain.BuildingName,
		CreatedAt:   	domain.CreatedAt,
		UpdatedAt:   	domain.UpdatedAt,
	}
}
