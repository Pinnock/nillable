package nillable

// The StringVal type represents values that are either nil or of type string.
type StringVal interface{}

// The String type wraps an underlying StringVal.
type String struct {
	val StringVal
}

// NewString returns a pointer to a new instance of String that wraps the
// underlying string, s.
func NewString(s string) *String {
	return &String{val: s}
}

// The Value method returns the underlying StringVal of this String.
func (s *String) Value() StringVal {
	return s.val
}

// The Set method sets the underlying StringVal to the given string, v.
func (s *String) Set(v string) {
	s.val = v
}

// The Clear method sets the underlying StringVal to nil.
func (s *String) Clear() {
	s.val = nil
}
