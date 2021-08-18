package response

import (
	"rentRoom/businesses/rents"
	"time"
)

type Rents struct {
	ID          int       `json:"id"`
	UserId    int	  `json:"user_id"`
	RoomId 	int	  `json:"room_id"`
	JumlahBayar   int       `json:"jumlah_pinjam"`
	TanggalPinjam   	string		`json:"tanggal_pinjam"`
	StatusPinjam   	bool		`json:"status_pinjam"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain rents.Domain) Rents {
	return Rents{
		ID:          	domain.ID,
		UserId:       domain.UserId,
		RoomId: 	domain.RoomId,
		JumlahBayar:      domain.JumlahBayar,
		TanggalPinjam: 	domain.TanggalPinjam,
		StatusPinjam: 	domain.StatusPinjam,
		CreatedAt:   	domain.CreatedAt,
		UpdatedAt:   	domain.UpdatedAt,
	}
}
