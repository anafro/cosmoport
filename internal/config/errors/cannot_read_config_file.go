package errors

import "fmt"

type ErrCannotReadConfigFile struct {
	error
	Path  string
	Cause error
}

func NewErrCannotReadConfigFile(path string, cause error) *ErrCannotReadConfigFile {
	return &ErrCannotReadConfigFile{
		Path:  path,
		Cause: cause,
	}
}

func (e *ErrCannotReadConfigFile) Error() string {
	return fmt.Sprintf(
		"cannot read config file '%s', because %s",
		e.Path,
		e.Cause.Error(),
	)
}
