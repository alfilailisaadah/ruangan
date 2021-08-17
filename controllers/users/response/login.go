package response

import (
	"rentRoom/businesses/users"
	"time"
)

type LoginResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	UserType  string    `json:"user_type"`
	Token string `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetLoginResponse(domain users.Domain, token string) LoginResponse {
	return LoginResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Username:  domain.Username,
		Password:  domain.Password,
		UserType:  domain.UserType,
		Token : 	token,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
