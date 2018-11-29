package repository

import (
	"database/sql"
	"log"

	"github.com/JuanTorr/project/model"
)

//UserDao user data access object
type UserDao struct {
	DB *sql.DB
}

//GetByEmail Gets the user information acording to the email given
func (dao UserDao) GetByEmail(email string) (u model.User, err error) {
	rows, err := dao.DB.Query(`
		SELECT
			test.user.user_id,
			test.user.email,
			test.user.password,
			test.user.name,
			test.user.surname,
			test.role.role_id,
			test.role.name,
			test.permission.permission_id,
			test.permission.name
		FROM test.user
		INNER JOIN test.role ON test.role.role_id = test.user.role_id
		INNER JOIN test.role_permission ON test.role_permission.role_id = test.role.role_id
		INNER JOIN test.permission ON test.permission.permission_id = test.role_permission.permission_id
		WHERE test.user.email = $1
	`, email)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := model.Permission{}
		err = rows.Scan(&u.ID, &u.Email, &u.Password, &u.Name, &u.Surname, &u.Role.ID, &u.Role.Name, &p.ID, &p.Name)
		if err != nil {
			return
		}
		u.Role.AddPermission(p)
	}
	return
}

//Register saves the user in the db
func (dao UserDao) Register(u model.User) (err error) {
	query := `
		INSERT INTO test.user (email, password, name, surname, role_id)
		VALUES ($1, $2, $3, $4, $5) RETURNING user_id`
	stmt, err := dao.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(u.Email, u.Password, u.Name, u.Surname, u.Role.ID)
	err = row.Scan(&u.ID)
	if err != nil {
		return
	}
	return
}
