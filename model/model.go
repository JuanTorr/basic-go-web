package model

//User struct
type User struct {
	ID       int
	Email    string
	Password string
	Name     string
	Surname  string
	Role     Role
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

//AddPermission appends the permissions
func (r *Role) AddPermission(p ...Permission) {
	r.Permissions = append(r.Permissions, p...)
}
