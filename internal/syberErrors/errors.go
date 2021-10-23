package syberErrors

import "fmt"

type ErrUnpackingJSON struct {
	Reason string
	Function string
}

func (e ErrUnpackingJSON) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

type ErrUserDoesntExist struct {
	Reason string
	Function string
}

func (e ErrUserDoesntExist) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}