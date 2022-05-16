package nillable

import "testing"

func TestDefaultBool(t *testing.T) {
	// Arrange
	var b Bool

	// Assert
	if b.val != nil {
		t.Fatalf("Expected nil, but got %v\n", b.val)
	}
}

func TestNewBoolWithoutError(t *testing.T) {
	// Arrange
	data := []string{
		"true", "True", "TRUE", "T", "t", "1",
		"false", "False", "FALSE", "F", "f", "0",
	}

	for idx, d := range data {
		// Act
		b, err := NewBool(d)

		// Assert
		if err != nil {
			t.Fatalf("Expected nil, but got %v\n", err)
		}
		if idx < 6 && b.val != true {
			t.Fatalf("Expected true, but got %v\n", b.val)
		}

		if idx > 5 && b.val != false {
			t.Fatalf("Expected false, but got %v\n", b.val)
		}
	}
}

func TestNewBoolWithError(t *testing.T) {
	// Arrange
	b, err := NewBool("tRUE")

	// Assert
	if b.val != nil {
		t.Fatalf("Expected val to be nil, but got %v\n", b.val)
	}

	if err == nil {
		t.Fatalf("Expected error to be non-nil but got %v\n", err)
	}
}

func TestBoolSet(t *testing.T) {
	// Arrange
	var a, b Bool

	// Act
	a.Set(true)
	b.Set(false)

	// Assert
	if a.val != true {
		t.Fatalf("Expected true but got %v\n", a.val)
	}
	if b.val != false {
		t.Fatalf("Expected false but got %v\n", b.val)
	}
}

func TestBoolValue(t *testing.T) {
	// Arrange
	var nval Bool
	tval, _ := NewBool("t")
	fval, _ := NewBool("f")

	// Assert
	if nval.Value() != nil {
		t.Fatalf("Expected nil, but got %v\n", nval.val)
	}
	if tval.Value() != true {
		t.Fatalf("Expected true, but got %v\n", tval.val)
	}
	if fval.Value() != false {
		t.Fatalf("Expected false, but got %v\n", tval.val)
	}

}

func TestBoolClear(t *testing.T) {
	// Arrange
	a, _ := NewBool("false")
	b, _ := NewBool("true")

	// Act
	a.Clear()
	b.Clear()

	// Assert
	if a.val != nil {
		t.Fatalf("Expected nil, but got %v\n", a.val)
	}
	if b.val != nil {
		t.Fatalf("Expected nil, but got %v\n", b.val)
	}
}
