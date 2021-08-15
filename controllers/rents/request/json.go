package request

import (
	"rentRoom/businesses/rents"
)


type Rents struct {
	JumlahBayar    int	`json:"jumlah_bayar"`
	TanggalPinjam    string	`json:"tanggal_pinjam"`
}

func (req *Rents) ToDomain() *rents.Domain {
	return &rents.Domain{
		JumlahBayar:   req.JumlahBayar,
		TanggalPinjam:   req.TanggalPinjam,
	}
}
