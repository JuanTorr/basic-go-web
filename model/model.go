package model

//User struct
type User struct {
	ID            int
	Email         string
	Password      string
	Name          string
	FirstSurname  string
	SecondSurname string
	Role          Role
}

//Role user's role
type Role struct {
	ID          int
	Name        string
	Permissions []Permission
}

//Permission system's permissions
type Permission struct {
	ID   int
	Name string
}
