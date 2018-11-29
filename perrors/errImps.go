package perrors

import "fmt"

func (err ErrRegNotFound) Error() string {
	return fmt.Sprintf("%s: %s", err.Message, err.Detail)
}

func (err ErrNoData) Error() string {
	return fmt.Sprintf("%s: %s", err.Message, err.Detail)
}

func (err ErrAuth) Error() string {
	return fmt.Sprintf("%s: %s", err.Message, err.Detail)
}

func (err ErrNotImplemented) Error() string {
	return fmt.Sprintf("%s: %s", err.Message, err.Detail)
}

func (err ErrIncorrectData) Error() string {
	return fmt.Sprintf("%s: %s", err.Message, err.Detail)
}
