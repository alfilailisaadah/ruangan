package response

import (
	"rentRoom/businesses/rents"
	"time"
)

type Rents struct {
	Id          int       `json:"id"`
	UserId    int	  `json:"user_id"`
	RoomId 	int	  `json:"room_id"`
	JumlahBayar   int       `json:"jumlah_pinjam"`
	TanggalPinjam   	string		`json:"tanggal_pinjam"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain rents.Domain) Rents {
	return Rents{
		Id:          	domain.ID,
		UserId:       domain.UserId,
		RoomId: 	domain.RoomId,
		JumlahBayar:      domain.JumlahBayar,
		TanggalPinjam: 	domain.TanggalPinjam,
		CreatedAt:   	domain.CreatedAt,
		UpdatedAt:   	domain.UpdatedAt,
	}
}
