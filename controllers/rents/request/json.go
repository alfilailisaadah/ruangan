package request

import (
	"rentRoom/businesses/rents"
)


type Rents struct {
	UserId    int	`json:"user_id"`
	RoomId    int	`json:"room_id"`
	JumlahBayar    int	`json:"jumlah_bayar"`
	TanggalPinjam    string	`json:"tanggal_pinjam"`
}

func (req *Rents) ToDomain() *rents.Domain {
	return &rents.Domain{
		JumlahBayar:   req.JumlahBayar,
		UserId:   req.UserId,
		RoomId:   req.RoomId,
		TanggalPinjam:   req.TanggalPinjam,
	}
}
