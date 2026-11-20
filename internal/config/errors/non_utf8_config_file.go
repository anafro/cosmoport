package errors

import "fmt"

type ErrNonUtf8ConfigFile struct {
	error
	Path string
}

func NewErrNonUtf8ConfigFile(path string) *ErrNonUtf8ConfigFile {
	return &ErrNonUtf8ConfigFile{
		Path: path,
	}
}

func (e *ErrNonUtf8ConfigFile) Error() string {
	return fmt.Sprintf(
		"config file '%s' is not utf-8",
		e.Path,
	)
}
