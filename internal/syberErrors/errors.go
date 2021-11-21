package syberErrors

import "fmt"

type ErrUnpackingJSON struct {
	Reason   string
	Function string
}

func (e ErrUnpackingJSON) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrUserDoesntExist struct {
	Reason   string
	Function string
}

func (e ErrUserDoesntExist) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrWrongPassword struct {
	Reason   string
	Function string
}

func (e ErrWrongPassword) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrUserExists struct {
	Reason   string
	Function string
}

func (e ErrUserExists) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrAuthorised struct {
	Reason   string
	Function string
}

func (e ErrAuthorised) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrNotLoggedin struct {
	Reason   string
	Function string
}

func (e ErrNotLoggedin) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrNotFeedNumber struct {
	Reason   string
	Function string
}

func (e ErrNotFeedNumber) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrInvalidEmail struct {
	Reason   string
	Function string
}

func (e ErrInvalidEmail) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrInvalidPassword struct {
	Reason   string
	Function string
}

func (e ErrInvalidPassword) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrInvalidLogin struct {
	Reason   string
	Function string
}

func (e ErrInvalidLogin) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrNoSession struct {
	Reason   string
	Function string
}

func (e ErrNoSession) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrDbError struct {
	Reason   string
	Function string
}

func (e ErrDbError) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrInternal struct {
	Reason   string
	Function string
}

func (e ErrInternal) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrValidate struct {
	Reason   string
	Function string
}

func (e ErrValidate) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrNoContent struct {
	Reason   string
	Function string
}

func (e ErrNoContent) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}

// -------------------------------------------------------------------------------

type ErrBadImage struct {
	Reason   string
	Function string
}

func (e ErrBadImage) Error() string {
	return fmt.Sprintf("error happend in %s, Reason: %s", e.Function, e.Reason)
}
