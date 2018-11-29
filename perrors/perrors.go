package perrors

//ErrMessage define los datos de un error
type ErrMessage struct {
	Message string
	Detail  string
}

//ErrRegNotFound defines the error when the there's no register i a query
type ErrRegNotFound struct {
	ErrMessage
}

//ErrNoData defines the error when the there's no data in a query or a map
type ErrNoData struct {
	ErrMessage
}

//ErrAuth defines the error when the user authentication has failed
type ErrAuth struct {
	ErrMessage
}

//ErrNotImplemented define el error cuando una funcionalidad no esta implementada
type ErrNotImplemented struct {
	ErrMessage
}

//ErrIncorrectData define el error cuando una funcionalidad no esta implementada
type ErrIncorrectData struct {
	ErrMessage
}
