package nillable

import "testing"

func TestDefaultString(t *testing.T) {
	// Arrange
	var s String

	// Assert
	if s.val != nil {
		t.Fatalf("Expected nil, but got %v\n", s.val)
	}
}

func TestNewString(t *testing.T) {
	// Arrange
	str := "hello"
	s := NewString(str)

	// Assert
	if s.val != str {
		t.Fatalf("Expected %q, but got %q\n", str, s.val)
	}
}

func TestStringSet(t *testing.T) {
	// Arrange
	var s String
	str := "Hello!"

	// Act
	s.Set(str)

	// Assert
	if s.val != str {
		t.Fatalf("Expected %q, but got %q\n", str, s.val)
	}
}

func TestStringClear(t *testing.T) {
	// Arrange
	s := NewString("Hello")

	// Act
	s.Clear()

	// Assert
	if s.val != nil {
		t.Fatalf("Expected nil, but got %q\n", s.val)
	}
}

func TestStringValue(t *testing.T) {
	// Arrange
	var str1 String
	str2 := NewString("")
	str3 := NewString("Hello, World!")

	if str1.Value() != nil {
		t.Fatalf("Expected nil, but got %v\n", str1.Value())
	}
	if str2.Value() != "" {
		t.Fatalf("Expected the empty string, but got %v\n", str2.Value())
	}
	if str3.Value() != "Hello, World!" {
		t.Fatalf("Expected \"Hello, World!\" , but got %q\n", str3.Value())
	}
}
