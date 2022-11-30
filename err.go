package errstack

import (
	"fmt"
	"runtime"
)

type Err struct {
	showMetadata bool
	filename     string
	line         int
	err          string
}

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

func (e *Err) SetShowMetadata(showMetadata bool) {
	e.showMetadata = showMetadata
}

func (e *Err) Error() string {
	if e.showMetadata == true {
		return fmt.Sprintf(`%s:%d="%v"`, e.filename, e.line, e.err)
	}

	return e.err
}
