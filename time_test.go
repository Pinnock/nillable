package nillable

import (
	"testing"
	"time"
)

func TestDefaultTime(t *testing.T) {
	// Arrange
	var tm Time

	// Assert
	if tm.val != nil {
		t.Fatalf("Expected nil, but got %v\n", tm.val)
	}
}

func TestNewTimeWithoutError(t *testing.T) {

	// Arrange
	const layout string = "Jan 2, 2006 3:04:05 PM"
	const value string = "Jun 5, 2022 6:49:23 PM"
	expval, _ := time.Parse(layout, value)

	// Act
	tm, err := NewTime(layout, value)

	// Assert
	if err != nil {
		t.Fatalf("Expected a nil error, but got %v\n", err)
	}
	if tm.val != expval {
		t.Fatalf("Expected value of %v, but got %v\n", expval, tm)
	}
}

func TestNewTimeWithError(t *testing.T) {
	// Arrange
	const layout string = "Jan 2, 2006 3:04:05 PM"
	const value string = "1/2/2006 3:04:05 PM"

	// Act
	tm, err := NewTime(layout, value)

	// Assert
	if err == nil {
		t.Fatalf("Expected non-nil error, but got %v\n", err)
	}
	if tm.val != nil {
		t.Fatalf("Expected nil, but got %v\n", tm.val)
	}
}

func TestTimeSet(t *testing.T) {
	// Arrange
	var tm1, tm2 Time
	const layout string = "Jan 2, 2006 3:04:05 PM"
	expval1, _ := time.Parse(layout, "Jun 22, 2022 5:36:00 PM")
	expval2, _ := time.Parse(layout, "Jun 22, 2022 5:36:00 AM")

	// Act
	tm1.Set(expval1)
	tm2.Set(expval2)

	// Assert
	if tm1.val != expval1 {
		t.Fatalf("Expected a value of %v, but got %v", expval1, tm1.val)
	}
	if tm2.val != expval2 {
		t.Fatalf("Expected a value of %v, but got %v", expval2, tm2.val)
	}
}

func TestTimeValue(t *testing.T) {
	// Arrange
	const layout string = "Jan 2, 2006 3:04:05 PM"
	const str1 string = "Jun 22, 2022 5:36:00 PM"
	const str2 string = "Jun 22, 2022 5:36:00PM"
	expval1, _ := time.Parse(layout, str1)

	// Act
	var tm0 Time
	tm1, _ := NewTime(layout, str1)
	tm2, _ := NewTime(layout, str2)

	// Assert
	if tm0.Value() != nil {
		t.Fatalf("Expected a value of nil, but got %v\n", tm0.Value())
	}
	if tm1.Value() != expval1 {
		t.Fatalf("Expected a value of %v, but got %v\n", expval1, tm1.Value())
	}
	if tm2.Value() != nil {
		t.Fatalf("Expected a value of %v, but got %v\n", nil, tm2.Value())
	}
}

func TestTimeClear(t *testing.T) {
	// Arrange
	const layout string = "Jan 2, 2006 3:04:05 PM"
	const str string = "Jun 22, 2022 5:36:00 PM"
	tm, _ := NewTime(layout, str)

	// Act
	tm.Clear()

	// Assert
	if tm.val != nil {
		t.Fatalf("Expected a value of %v, but got %v\n", nil, tm.val)
	}
}
