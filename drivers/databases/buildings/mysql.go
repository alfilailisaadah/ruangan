package buildings

import (
	"context"
	"rentRoom/businesses/buildings"

	"gorm.io/gorm"
)

type mysqlBuildingsRepository struct {
	Conn *gorm.DB
}

func NewMySQLBuidingsRepository(conn *gorm.DB) buildings.Repository {
	return &mysqlBuildingsRepository{
		Conn: conn,
	}
}

func (nr *mysqlBuildingsRepository) Fetch(ctx context.Context, page, perpage int) ([]buildings.Domain, int, error) {
	rec := []Buildings{}

	offset := (page - 1) * perpage
	err := nr.Conn.Preload("Room").Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []buildings.Domain{}, 0, err
	}

	var totalData int64
	err = nr.Conn.Count(&totalData).Error
	if err != nil {
		return []buildings.Domain{}, 0, err
	}

	var domainBuildings []buildings.Domain
	for _, value := range rec {
		domainBuildings = append(domainBuildings, value.toDomain())
	}
	return domainBuildings, int(totalData), nil
}

func (nr *mysqlBuildingsRepository) GetByID(ctx context.Context, buildingsId int) (buildings.Domain, error) {
	rec := Buildings{}
	err := nr.Conn.Where("id = ?", buildingsId).First(&rec).Error
	if err != nil {
		return buildings.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlBuildingsRepository) GetByTitle(ctx context.Context, buildingsTitle string) (buildings.Domain, error) {
	rec := Buildings{}
	err := nr.Conn.Where("title = ?", buildingsTitle).First(&rec).Error
	if err != nil {
		return buildings.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlBuildingsRepository) Store(ctx context.Context, buildingsDomain *buildings.Domain) (buildings.Domain, error) {
	rec := fromDomain(buildingsDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return buildings.Domain{}, result.Error
	}

	err := nr.Conn.Preload("Room").First(&rec, rec.Id).Error
	if err != nil {
		return buildings.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlBuildingsRepository) Update(ctx context.Context, buildingsDomain *buildings.Domain) (buildings.Domain, error) {
	rec := fromDomain(buildingsDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return buildings.Domain{}, result.Error
	}

	err := nr.Conn.Preload("Room").First(&rec, rec.Id).Error
	if err != nil {
		return buildings.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}
