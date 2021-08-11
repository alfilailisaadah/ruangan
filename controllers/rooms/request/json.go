package request

import "rentRoom/businesses/rooms"

type Rooms struct {
	StatusPinjam    string `json:"statusPinjam"`
	HargaPinjam int    `json:"hargaPinjam"`
}

func (req *Rooms) ToDomain() *rooms.Domain {
	return &rooms.Domain{
		StatusPinjam:    req.StatusPinjam,
		HargaPinjam: req.HargaPinjam,
	}
}
