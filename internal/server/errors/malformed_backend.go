package errors

import "fmt"

type ErrMalformedBackend struct {
	error
	Backend string
	Message string
}

func NewErrMalformedBackend(backend, message string) *ErrMalformedBackend {
	return &ErrMalformedBackend{
		Backend: backend,
		Message: message,
	}
}

func (e *ErrMalformedBackend) Error() string {
	return fmt.Sprintf(
		"the backend '%s' is malformed: %s",
		e.Backend,
		e.Message,
	)
}
