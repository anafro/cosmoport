package errors

import "fmt"

type ErrMalformedDestination struct {
	error
}

func NewErrMalformedDestination() *ErrMalformedDestination {
	return &ErrMalformedDestination{}
}

func (e *ErrMalformedDestination) Error() string {
	return fmt.Sprintf(
		"neither `backend` or `static` is specified for some rule in the config",
	)
}
