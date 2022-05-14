package nillable

import "strconv"

// The FloatVal type represents values that are either nil or of type float64.
type FloatVal interface{}

// The Float type wraps an underlying FloatVal.
type Float struct {
	val FloatVal
}

// NewFloat returns a pointer to a new instance of Float that wraps the
// underlying FloatVal obtained by parsing the given string, s. If the string
// was not successfully parsed then a non-nil error is also returned. Otherwsie,
// the error is nil.
func NewFloat(s string) (*Float, error) {
	return parseFloat(s)
}

// The Value method returns the underlying FloatVal of this Float.
func (f *Float) Value() FloatVal {
	return f.val
}

// The Set method sets the underlying FloatVal to the given float64, v.
func (f *Float) Set(v float64) {
	f.val = v
}

// The Clear method sets the underlying FloatVal to nil.
func (f *Float) Clear() {
	f.val = nil
}

func parseFloat(s string) (*Float, error) {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return &Float{val: nil}, err
	}
	return &Float{val: val}, nil
}
