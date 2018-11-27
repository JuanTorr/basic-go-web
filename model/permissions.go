package model

//PermissionConst constantes de permisos
type PermissionConst int

const (
	//EditUsername permission to edit the user's name
	EditUsername PermissionConst = 1
	//Editemail permission to edit the user's email
	Editemail PermissionConst = 2
)
