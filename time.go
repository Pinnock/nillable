package nillable

import (
	"time"
)

// The TimeVal type represents values that are either nil or of type time.Time
type TimeVal interface{}

// The Time type wraps an underlying TimeVal
type Time struct {
	val TimeVal
}

// NewTime returns a pointer to a new instance of Time that wraps the underlying
// TimeVal obtained by parsing the given string, value, based on the same rules
// of time.Parse from the standard library. If the string was not successfully,
// parsed then a non-nil error is also returned. Otherwsie the error is nil.
func NewTime(layout, value string) (*Time, error) {
	return parseTime(layout, value)
}

// The Value() method returns the underlying TimeVal of this Time.
func (t *Time) Value() TimeVal {
	return t.val
}

// The Set method sets the underlying TimeVal to the given time.Time, v.
func (t *Time) Set(v time.Time) {
	t.val = v
}

// The Clear mehtod sets the underlying TimeVal to nil
func (t *Time) Clear() {
	t.val = nil
}

func parseTime(layout, value string) (*Time, error) {
	val, err := time.Parse(layout, value)
	if err != nil {
		return &Time{val: nil}, err
	}
	return &Time{val: val}, nil
}
