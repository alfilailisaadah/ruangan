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
type usersRepository struct {
	conn *gorm.DB
}

func (nr *rentsRepository) Store(ctx context.Context, rentsDomain *rents.Domain) (rents.Domain, error) {
	rec := fromDomain(rentsDomain)

	result := nr.conn.Create(&rec)
	if result.Error != nil {
		return rents.Domain{}, result.Error
	}

	err := nr.conn.Preload("User","Room").First(&rec, rec.ID).Error
	if err != nil {
		return rents.Domain{}, result.Error
	}
	return rec.ToDomain(), nil
}

func (nr *rentsRepository) GetById(ctx context.Context,userId int) (rents.Domain, error) {
	rec := Rents{}
	err := nr.conn.Where("user_id = ?", userId).First(&rec).Error
	if err != nil {
		return rents.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (nr *rentsRepository) Fetch(ctx context.Context, page, perpage int) ([]rents.Domain, int, error) {
	rec := []Rents{}

	offset := (page - 1) * perpage
	err := nr.conn.Preload("Room").Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []rents.Domain{}, 0, err
	}

	var totalData int64
	err = nr.conn.Count(&totalData).Error
	if err != nil {
		return []rents.Domain{}, 0, err
	}

	var domainRents []rents.Domain
	for _, value := range rec {
		domainRents = append(domainRents, value.ToDomain())
	}
	return domainRents, int(totalData), nil
}

func (cr *rentsRepository) Find(ctx context.Context, rentStatus string) ([]rents.Domain, error) {
	rec := []Rents{}

	query := cr.conn

	if rentStatus != "" {
		query = query.Where("rentStatus = ?", rentStatus)
	}

	err := query.Find(&rec).Error
	if err != nil {
		return []rents.Domain{}, err
	}

	categoryDomain := []rents.Domain{}
	for _, value := range rec {
		categoryDomain = append(categoryDomain, value.ToDomain())
	}

	return categoryDomain, nil
}

