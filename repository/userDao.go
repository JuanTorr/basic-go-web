package repository

import (
	"math/rand"

	"github.com/JuanTorr/project/bberrors"
	"github.com/JuanTorr/project/model"
)

var db map[string]interface{}

//UserDao user data access object
type UserDao struct{}

func init() {
	db = make(map[string]interface{})
	db["users"] = []model.User{
		model.User{
			ID:           rand.Intn(999999) + 1,
			Email:        "email.com",
			Name:         "Jose Alfredo",
			FirstSurname: "Jiménez",
			Password:     "1",
			Role: model.Role{
				ID:   1,
				Name: "Admin",
				Permissions: []model.Permission{
					model.Permission{ID: 1, Name: "editusername"},
					model.Permission{ID: 2, Name: "editemail"},
				},
			},
		},
	}
}

//GetByEmail Gets the user information acording to the email given
func (dao UserDao) GetByEmail(email string) (model.User, error) {
	users := db["users"].([]model.User)
	for _, u := range users {
		if u.Email == email {
			return u, nil
		}
	}
	return model.User{}, bberrors.ErrRegNotFound
}

//Register saves the user in the db
func (dao UserDao) Register(u model.User) (err error) {
	u.ID = rand.Intn(999999) + 1
	db["users"] = append(db["users"].([]model.User), u)
	return
}
