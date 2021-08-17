package rents

import (
	"context"
	"fmt"
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
	errr := nr.conn.Preload("Room").First(&rec, rec.ID).Error
	if errr != nil {
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

func (nr *rentsRepository) GetByRentId(ctx context.Context,rentId int) (rents.Domain, error) {
	rec := Rents{}
	err := nr.conn.Where("id = ?", rentId).First(&rec).Error
	if err != nil {
		return rents.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (nr *rentsRepository) GetRoomById(ctx context.Context,roomId int) (rents.Domain, error) {
	rec := Rents{}
	err := nr.conn.Where("room_id = ?", roomId).First(&rec).Error
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

func (nr *rentsRepository) Update(ctx context.Context, rentsDomain *rents.Domain) (rents.Domain, error) {
	rec := fromDomain(rentsDomain)
	fmt.Println(rec)
	result := nr.conn.Updates(&rec)
	if result.Error != nil {
		return rents.Domain{}, result.Error
	}

	err := nr.conn.Preload("Rents").First(&rec,rec.ID).Error
	if err != nil {
		return rents.Domain{}, result.Error
	}


	return rec.ToDomain(), nil
}


