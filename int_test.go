package nillable

import "testing"

func TestDefaultInt(t *testing.T) {
	// Arrange
	var i Int

	// Assert
	if i.val != nil {
		t.Fatalf("Expected nil, but got %v\n", i.val)
	}
}

func TestNewIntWithoutError(t *testing.T) {
	// Arrange
	i, err := NewInt("10")

	// Assert
	if err != nil {
		t.Fatalf("Expected a nil error, but got %v\n", err)
	}
	if i.val != 10 {
		t.Fatalf("Expected value of 10, but got %v\n", i.val)
	}
}

func TestNewIntWithError(t *testing.T) {
	// Arrange
	i, err := NewInt("10.1")

	// Assert
	if err == nil {
		t.Fatalf("Expected a non-nil error, but got %v\n", err)
	}
	if i.val != nil {
		t.Fatalf("Expected value of nil, but got %v\n", i.val)
	}
}
func TestIntSet(t *testing.T) {
	// Arrange
	var a, b Int

	// Act
	a.Set(-1)
	b.Set(15)

	// Assert
	if a.val != -1 {
		t.Fatalf("Expected a value of -1, but got %v\n", a.val)
	}
	if b.val != 15 {
		t.Fatalf("Expected a value of 15, but got %v\n", b.val)
	}
}

func TestIntClear(t *testing.T) {
	// Arrange
	i, _ := NewInt("10")

	// Act
	i.Clear()

	// Assert
	if i.val != nil {
		t.Fatalf("Expected a value of nil, but got %v\n", i.val)
	}
}
