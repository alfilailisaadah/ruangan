package rents

import (
	"context"
	"rentRoom/businesses/rents"

	"gorm.io/gorm"
)

type rentsRepository struct {
	conn *gorm.DB
}

func NewRentsRepository(conn *gorm.DB) rents.Repository {
	return &rentsRepository{
		conn: conn,
	}
}

func (nr *rentsRepository) Store(ctx context.Context, rentsDomain *rents.Domain) (rents.Domain, error) {
	rec := fromDomain(rentsDomain)

	result := nr.conn.Create(&rec)
	if result.Error != nil {
		return rents.Domain{}, result.Error
	}

	err := nr.conn.Preload("User").First(&rec, rec.ID).Error
	if err != nil {
		return rents.Domain{}, result.Error
	}
	return rec.ToDomain(), nil
}

