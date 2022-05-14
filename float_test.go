package nillable

import "testing"

func TestDefaultFloat(t *testing.T) {
	// Arrange
	var f Float

	// Assert
	if f.val != nil {
		t.Fatalf("Expected value of nil, but got %v\n", f.val)
	}
}

func TestNewFloatWithoutError(t *testing.T) {
	// Arrange
	f1, err1 := NewFloat("10")
	f2, err2 := NewFloat("-1.2")

	// Assert
	if err1 != nil {
		t.Fatalf("Expected a nil error, but got %v\n", err1)
	}
	if f1.val != 10.0 {
		t.Fatalf("Expected value of 10, but got %v\n", f1.val)
	}
	if err2 != nil {
		t.Fatalf("Expected a nil error, but got %v\n", err2)
	}
	if f2.val != -1.2 {
		t.Fatalf("Expected value of 10, but got %v\n", f2.val)
	}
}

func TestNewFloatWithError(t *testing.T) {
	// Arrange
	f, err := NewFloat("1O.1")

	// Assert
	if err == nil {
		t.Fatalf("Expected a non-nil error, but got %v\n", err)
	}
	if f.val != nil {
		t.Fatalf("Expected value of nil, but got %v\n", f.val)
	}
}

func TestFloatSet(t *testing.T) {
	// Arrange
	var a, b Float

	// Act
	a.Set(-1)
	b.Set(15.123)

	// Assert
	if a.val != -1.0 {
		t.Fatalf("Expected a value of -1, but got %v\n", a.val)
	}
	if b.val != 15.123 {
		t.Fatalf("Expected a value of 15.123, but got %v\n", b.val)
	}
}

func TestFloatValue(t *testing.T) {
	// Arrange
	var f1 Float
	f2, _ := NewFloat("-1")
	f3, _ := NewFloat("1.O")

	// Assert
	if f1.Value() != nil {
		t.Fatalf("Expected a value of nil, but got %v\n", f1.val)
	}
	if f2.Value() != -1.0 {
		t.Fatalf("Expected a value of -1.0, but got %v\n", f2.val)
	}
	if f3.Value() != nil {
		t.Fatalf("Expected a value of nil, but got %v\n", f1.val)
	}
}

func TestFloatClear(t *testing.T) {
	// Arrange
	f, _ := NewFloat("10")

	// Act
	f.Clear()

	// Assert
	if f.val != nil {
		t.Fatalf("Expected a value of nil, but got %v\n", f.val)
	}
}
