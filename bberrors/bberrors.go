package bberrors

import (
	"errors"
)

//ErrRegNotFound s
var ErrRegNotFound = errors.New("Register not found")

//ErrAuth defines the error when the user authentication has failed
var ErrAuth = errors.New("User authentication has failed")
