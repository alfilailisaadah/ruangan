package request

import "rentRoom/businesses/buildings"

type Buildings struct {
	BuildingName    string	`json:"building_name"`
	BuildingAddr    string	`json:"building_addr"`
}

func (req *Buildings) ToDomain() *buildings.Domain {
	return &buildings.Domain{
		BuildingName:   req.BuildingName,
		BuildingAddr:   req.BuildingAddr,
	}
}
