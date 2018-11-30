package service

import "github.com/JuanTorr/basic-go-web/model"

//UserDaoService service to get user's information
type UserDaoService interface {
	GetByEmail(string) (model.User, error)
	Register(*model.User) error
}
