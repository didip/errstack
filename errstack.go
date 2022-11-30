package errstack

import "strings"

type ErrStack struct {
	stack        []*Err
	showMetadata bool
}

func NewString(err string) *ErrStack {
	es := &ErrStack{
		stack:        make([]*Err, 0),
		showMetadata: true,
	}
	es.Append(err)

	return es
}

func (es *ErrStack) AppendErr(err *Err) {
	es.stack = append(es.stack, err)
}

func (es *ErrStack) Append(err string) {
	es.AppendErr(NewErr(err))
}

func (es *ErrStack) SetShowMetadata(showMetadata bool) {
	es.showMetadata = showMetadata
}

func (es *ErrStack) GetAll() []*Err {
	stack := make([]*Err, len(es.stack))
	copy(stack, es.stack)

	// Reverse the stack output to make it LIFO
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	return stack
}

func (es *ErrStack) PopAll() []*Err {
	stack := es.GetAll()
	es.stack = make([]*Err, 0)
	return stack
}

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
