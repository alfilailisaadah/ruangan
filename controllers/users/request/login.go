package request

import "rentRoom/businesses/users"

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *UserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Username: req.Username,
		Password: req.Password,
	}
}