package request

import (
	"rentRoom/businesses/rents"
)


type Rents struct {
	ID		int `json:"id"`
	UserId    int	`json:"user_id"`
	RoomId    int	`json:"room_id"`
	JumlahBayar    int	`json:"jumlah_bayar"`
	StatusPinjam    bool	`json:"status_pinjam"`
	TanggalPinjam    string	`json:"tanggal_pinjam"`
}

func (req *Rents) ToDomain() *rents.Domain {
	return &rents.Domain{
		ID: req.ID,
		JumlahBayar:   req.JumlahBayar,
		UserId:   req.UserId,
		RoomId:   req.RoomId,
		StatusPinjam:   req.StatusPinjam,
		TanggalPinjam:   req.TanggalPinjam,
	}
}
