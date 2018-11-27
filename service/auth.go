package service

import (
	"github.com/JuanTorr/project/bberrors"
	"github.com/JuanTorr/project/model"
)

//AuthenticateUser validate if the given email and password match
func AuthenticateUser(
	userDao UserDaoService, email, password string,
) (u model.User, err error) {
	u, err = userDao.GetByEmail(email)
	if err != nil {
		return
	}
	if u.Password != password {
		return u, bberrors.ErrAuth
	}
	return u, nil
}

//RegisterUser registers the given user
func RegisterUser(userDao UserDaoService, u model.User) error {
	return userDao.Register(u)
}
