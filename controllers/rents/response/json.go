package response

import (
	"rentRoom/businesses/rents"
	"time"
)

type Rooms struct {
	Id          int       `json:"id"`
	UserId    int	  `json:"user_id"`
	RoomsId 	int	  `json:"rooms_id"`
	JumlahBayar   int       `json:"jumlah_pinjam"`
	TanggalPinjam   	string		`json:"tanggal_pinjam"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain rents.Domain) Rooms {
	return Rooms{
		Id:          	domain.ID,
		UserId:       domain.UserId,
		RoomsId: 	domain.RoomsId,
		JumlahBayar:      domain.JumlahBayar,
		TanggalPinjam: 	domain.TanggalPinjam,
		CreatedAt:   	domain.CreatedAt,
		UpdatedAt:   	domain.UpdatedAt,
	}
}
