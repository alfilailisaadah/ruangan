package buildings

import (
	buildingsUsecase "rentRoom/businesses/buildings"
	"time"
)

type Buildings struct {
	Id         int
	BuildingName      string
	BuildingAddr    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func fromDomain(domain *buildingsUsecase.Domain) *Buildings {
	return &Buildings{
		Id:         	domain.ID,
		BuildingName:   domain.BuildingName,
		BuildingAddr:   domain.BuildingAddr,
	}
}

func (rec *Buildings) toDomain() buildingsUsecase.Domain {
	return buildingsUsecase.Domain{
		ID:           rec.Id,
		BuildingName: rec.BuildingName,
		BuildingAddr: rec.BuildingAddr,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
	}
}
