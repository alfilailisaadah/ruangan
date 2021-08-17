package response

import (
	"rentRoom/businesses/buildings"
	"time"
)

type Buildings struct {
	Id           int       `json:"id"`
	BuildingName    string	`json:"building_name"`
	BuildingAddr    string	`json:"building_addr"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromDomain(domain buildings.Domain) Buildings {
	return Buildings{
		Id:           domain.ID,
		BuildingName:        domain.BuildingName,
		BuildingAddr:      domain.BuildingAddr,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
