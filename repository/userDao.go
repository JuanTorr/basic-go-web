package repository

import (
	"database/sql"
	"fmt"

	"github.com/JuanTorr/basic-go-web/model"
	"github.com/JuanTorr/basic-go-web/perrors"
)

//UserDao user data access object
type UserDao struct {
	DB *sql.DB
}

//GetByEmail Gets the user information acording to the email given
func (dao UserDao) GetByEmail(email string) (u model.User, err error) {
	rows, err := dao.DB.Query(`
		SELECT
			user.tus_id,
			user.tus_email,
			user.tus_password,
			user.tus_name,
			user.tus_surname,
			role.tro_id,
			role.tro_name,
			permission.tpe_id,
			permission.tpe_name
		FROM tuser user
		INNER JOIN trole role ON role.tro_id = user.tus_role_id
		INNER JOIN rrol_permission role_permission ON role_permission.rrp_role_id = role.tro_id
		INNER JOIN tpermission permission ON permission.tpe_id = role_permission.rrp_permission_id
		WHERE user.tus_email = ?
	`, email)
	if err != nil {
		return
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		count++
		p := model.Permission{}
		err = rows.Scan(&u.ID, &u.Email, &u.Password, &u.Name, &u.Surname, &u.Role.ID, &u.Role.Name, &p.ID, &p.Name)
		if err != nil {
			return
		}
		u.Role.AddPermission(p)
	}
	if count == 0 {
		err = perrors.NewErrRegNotFound("User not found", email)
	}
	return
}

//Register saves the user in the db
func (dao UserDao) Register(u *model.User) (err error) {
	query := `INSERT INTO
		tuser (tus_email, tus_password, tus_name, tus_surname, tus_role_id)
		VALUES (?,?,?,?,?)`
	result, err := dao.DB.Exec(query, u.Email, u.Password, u.Name, u.Surname, u.Role.ID)
	if err != nil {
		return
	}
	v, err := result.LastInsertId()
	if err != nil {
		return
	}
	(*u).ID = int(v)
	fmt.Printf("\n*********************\n%+v\n", u)
	return
}
