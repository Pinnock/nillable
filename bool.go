package nillable

import (
	"strconv"
)

// The BoolVal type represents values that are either nil or of type bool.
type BoolVal interface{}

// The Bool type wraps an underlying BoolVal
type Bool struct {
	val BoolVal
}

// NewBool returns a pointer to a new instance of Bool that wraps the underlying
// BoolVal obtained by parsing the given string, s. If the string was not
// successfully parsed then a non-nil error is also returned. Otherwsie the
// error is nil. The parsable strings are: "1", "t", "T", "true", "TRUE",
// "True", "0", "f", "F", "false", "FALSE", and "False".
func NewBool(s string) (*Bool, error) {
	return parseBool(s)
}

// The Value method returns the underlying BoolVal of this Bool.
func (b *Bool) Value() BoolVal {
	return b.val
}

// The Set method sets the underlying BoolVal to the given bool, v.
func (b *Bool) Set(v bool) {
	b.val = v
}

// The Clear method sets the underlying BoolVal to nil.
func (b *Bool) Clear() {
	b.val = nil
}

func parseBool(s string) (*Bool, error) {
	val, err := strconv.ParseBool(s)
	if err != nil {
		return &Bool{val: nil}, err
	}
	return &Bool{val: val}, nil
}
