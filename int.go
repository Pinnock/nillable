package nillable

import (
	"strconv"
)

// The IntVal type represents values that are either nil or of type int.
type IntVal interface{}

// The Int type wraps an underlying IntVal
type Int struct {
	val IntVal
}

// NewInt returns a pointer to a new instance of Int that wraps the underlying
// IntVal obtained by parsing the given string, s. If the string was not
// successfully parsed then a non-nil error is also returned. Otherwsie the
// error is nil.
func NewInt(s string) (*Int, error) {
	return parseInt(s)
}

// The Value method returns the underlying IntVal of this Int.
func (i *Int) Value() IntVal {
	return i.val
}

// The Set method sets the underlying IntVal to the given int, v.
func (i *Int) Set(v int) {
	i.val = v
}

// The Clear method sets the underlying IntVal to nil.
func (i *Int) Clear() {
	i.val = nil
}

func parseInt(s string) (*Int, error) {
	val, err := strconv.Atoi(s)
	if err != nil {
		return &Int{val: nil}, err
	}
	return &Int{val: val}, nil
}
