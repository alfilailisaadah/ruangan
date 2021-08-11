package rooms

import (
	"context"
	"rentRoom/businesses/rooms"

	"gorm.io/gorm"
)

type mysqlRoomsRepository struct {
	Conn *gorm.DB
}

func NewMySQLRoomsRepository(conn *gorm.DB) rooms.Repository {
	return &mysqlRoomsRepository{
		Conn: conn,
	}
}

func (nr *mysqlRoomsRepository) Fetch(ctx context.Context, page, perpage int) ([]rooms.Domain, int, error) {
	rec := []Rooms{}

	offset := (page - 1) * perpage
	err := nr.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []rooms.Domain{}, 0, err
	}

	var totalData int64
	err = nr.Conn.Count(&totalData).Error
	if err != nil {
		return []rooms.Domain{}, 0, err
	}

	var domainRooms []rooms.Domain
	for _, value := range rec {
		domainRooms = append(domainRooms, value.toDomain())
	}
	return domainRooms, int(totalData), nil
}

func (nr *mysqlRoomsRepository) GetByID(ctx context.Context, roomsId int) (rooms.Domain, error) {
	rec := Rooms{}
	return rec.toDomain(), nil
}

func (nr *mysqlRoomsRepository) GetByTitle(ctx context.Context, roomsTitle string) (rooms.Domain, error) {
	rec := Rooms{}
	err := nr.Conn.Where("title = ?", roomsTitle).First(&rec).Error
	if err != nil {
		return rooms.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlRoomsRepository) Store(ctx context.Context, roomsDomain *rooms.Domain) error {
	rec := fromDomain(roomsDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
