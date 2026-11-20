package errors

import "fmt"

type ErrMalformedConfigFile struct {
	error
	Path  string
	Cause error
}

func NewMalformedConfigFile(path string, cause error) *ErrMalformedConfigFile {
	return &ErrMalformedConfigFile{
		Path:  path,
		Cause: cause,
	}
}

func (e *ErrMalformedConfigFile) Error() string {
	return fmt.Sprintf(
		"cannot read config file '%s', because %s",
		e.Path,
		e.Cause.Error(),
	)
}
