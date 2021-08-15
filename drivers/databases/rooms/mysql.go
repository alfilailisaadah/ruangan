package rooms

import (
	"context"
	"rentRoom/businesses/rooms"

	"gorm.io/gorm"
)

type roomsRepository struct {
	conn *gorm.DB
}

func NewRoomsRepository(conn *gorm.DB) rooms.Repository {
	return &roomsRepository{
		conn: conn,
	}
}

func (nr *roomsRepository) Store(ctx context.Context, newsDomain *rooms.Domain) (rooms.Domain, error) {
	rec := fromDomain(newsDomain)

	result := nr.conn.Create(&rec)
	if result.Error != nil {
		return rooms.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}

func (cr *roomsRepository) Find(ctx context.Context, rentStatus string) ([]rooms.Domain, error) {
	rec := []Rooms{}

	query := cr.conn

	if rentStatus != "" {
		query = query.Where("rentStatus = ?", rentStatus)
	}

	err := query.Find(&rec).Error
	if err != nil {
		return []rooms.Domain{}, err
	}

	categoryDomain := []rooms.Domain{}
	for _, value := range rec {
		categoryDomain = append(categoryDomain, value.ToDomain())
	}

	return categoryDomain, nil
}

func (cr *roomsRepository) FindByID(id int) (rooms.Domain, error) {
	rec := Rooms{}

	if err := cr.conn.Where("id = ?", id).First(&rec).Error; err != nil {
		return rooms.Domain{}, err
	}
	return rec.ToDomain(), nil
}
