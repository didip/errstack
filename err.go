package errstack

import (
	"fmt"
	"runtime"
)

// Err struct contains filename, line number, and error message
type Err struct {
	showMetadata bool
	filename     string
	line         int
	err          string
}

// NewErr returns *Err struct
func NewErr(err string) *Err {
	_, filename, line, _ := runtime.Caller(2)
	e := &Err{
		filename:     filename,
		line:         line,
		err:          err,
		showMetadata: true,
	}
	return e
}

// SetShowMetadata is a flag to display/hide filename and line number
func (e *Err) SetShowMetadata(showMetadata bool) {
	e.showMetadata = showMetadata
}

// Error satisfies the error interface
func (e *Err) Error() string {
	if e.showMetadata {
		return fmt.Sprintf(`%s:%d="%v"`, e.filename, e.line, e.err)
	}

	return e.err
}
