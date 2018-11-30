package bgwerrors

//NewErrRegNotFound contructor del error NewErrRegNotFound
func NewErrRegNotFound(s ...string) (err ErrRegNotFound) {
	err.Message = s[0]
	if len(s) > 1 {
		err.Detail = s[1]
	}
	return
}

//NewErrNoData contructor del error NewErrNoData
func NewErrNoData(s ...string) (err ErrNoData) {
	err.Message = s[0]
	if len(s) > 1 {
		err.Detail = s[1]
	}
	return
}

//NewErrAuth contructor del error NewErrAuth
func NewErrAuth(s ...string) (err ErrAuth) {
	err.Message = s[0]
	if len(s) > 1 {
		err.Detail = s[1]
	}
	return
}

//NewErrNotImplemented contructor del error NewErrNotImplemented
func NewErrNotImplemented(s ...string) (err ErrNotImplemented) {
	err.Message = s[0]
	if len(s) > 1 {
		err.Detail = s[1]
	}
	return
}

//NewErrIncorrectData contructor del error ErrIncorrectData
func NewErrIncorrectData(s ...string) (err ErrNotImplemented) {
	err.Message = s[0]
	if len(s) > 1 {
		err.Detail = s[1]
	}
	return
}
