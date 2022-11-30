package errstack

import (
	"strings"
	"sync"
)

// ErrStack contains the error stack
type ErrStack struct {
	stack        []*Err
	showMetadata bool
	mtx          *sync.Mutex
}

// New creates a new ErrStack given an error string
func New(err string) *ErrStack {
	es := &ErrStack{
		stack:        make([]*Err, 0),
		showMetadata: true,
		mtx:          &sync.Mutex{},
	}
	es.Append(err)

	return es
}

// AppendErr allows users to append *Err struct into the stack
func (es *ErrStack) AppendErr(err *Err) {
	es.mtx.Lock()
	es.stack = append(es.stack, err)
	es.mtx.Unlock()
}

// Append allows users to append an error string
func (es *ErrStack) Append(err string) {
	es.AppendErr(NewErr(err))
}

// SetShowMetadata is a flag to display/hide filename and line number
func (es *ErrStack) SetShowMetadata(showMetadata bool) {
	es.mtx.Lock()
	es.showMetadata = showMetadata
	es.mtx.Unlock()
}

// GetAll returns a list of *Err structs in a LIFO fashion
func (es *ErrStack) GetAll() []*Err {
	stack := make([]*Err, len(es.stack))
	copy(stack, es.stack)

	// Reverse the stack output to make it LIFO
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	return stack
}

// PopAll returns a list of *Err structs in a LIFO fashion and clears the stack.
func (es *ErrStack) PopAll() []*Err {
	stack := es.GetAll()

	es.mtx.Lock()
	es.stack = make([]*Err, 0)
	es.mtx.Unlock()
	return stack
}

// Error satisfies the error interface
func (es *ErrStack) Error() string {
	asString := make([]string, len(es.stack))

	for i, err := range es.GetAll() {
		err.SetShowMetadata(es.showMetadata)
		asString[i] = err.Error()
	}

	if es.showMetadata {
		return strings.Join(asString, " ")
	}

	return strings.Join(asString, ", ")
}
