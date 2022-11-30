package errstack

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
)

// Err struct contains filename, line number, and error message
type Err struct {
	showMetadata bool
	trimFilename bool
	filename     string
	line         int
	err          string
	mtx          *sync.Mutex
}

// NewErr returns *Err struct
func NewErr(err string) *Err {
	_, filename, line, _ := runtime.Caller(2)
	e := &Err{
		filename:     filename,
		line:         line,
		err:          err,
		showMetadata: true,
		trimFilename: false,
		mtx:          &sync.Mutex{},
	}
	return e
}

// getFilename is helper function that returns shorten or original filename
func (e *Err) getFilename() string {
	if e.trimFilename {
		chunks := strings.Split(e.filename, string(os.PathSeparator))
		chunksLength := len(chunks)

		domainIndex := -1

		for i, chunk := range chunks {
			if strings.HasSuffix(chunk, ".com") || strings.HasSuffix(chunk, ".edu") || strings.HasSuffix(chunk, ".org") || strings.HasSuffix(chunk, ".net") || strings.HasSuffix(chunk, ".io") {
				domainIndex = i
				break
			}
		}

		return strings.Join(chunks[domainIndex:chunksLength], string(os.PathSeparator))
	}
	return e.filename
}

// SetShowMetadata is a flag to display/hide filename and line number
func (e *Err) SetShowMetadata(showMetadata bool) *Err {
	e.mtx.Lock()
	e.showMetadata = showMetadata
	e.mtx.Unlock()

	return e
}

// SetTrimFilename is a flag to trim filename
func (e *Err) SetTrimFilename(trimFilename bool) *Err {
	e.mtx.Lock()
	e.trimFilename = trimFilename
	e.mtx.Unlock()

	return e
}

// Error satisfies the error interface
func (e *Err) Error() string {
	if e.showMetadata {
		return fmt.Sprintf(`%s:%d="%v"`, e.getFilename(), e.line, e.err)
	}

	return e.err
}
