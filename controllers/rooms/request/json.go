package request

import "rentRoom/businesses/rooms"

type Rooms struct {
	RoomName    string	`json:"rooms_name"`
	RentStatus    bool	`json:"rent_status"`
	RentPrice   	int		`json:"rent_price"`
	BuildingId   	int		`json:"building_id"`
}

func (req *Rooms) ToDomain() *rooms.Domain {
	return &rooms.Domain{
		RoomName:   req.RoomName,
		RentStatus:   req.RentStatus,
		RentPrice: 		req.RentPrice,
		BuildingId: req.BuildingId,
	}
}
