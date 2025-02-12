package internal

import (
	"testing"
)

// TestAddition tests the addition of two numbers
func TestAddition(t *testing.T) {
	result := 2 + 2
	expected := 4
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// TestSubtraction tests the subtraction of two numbers
func TestSubtraction(t *testing.T) {
	result := 10 - 5
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// TestMultiplication tests the multiplication of two numbers
func TestMultiplication(t *testing.T) {
	result := 3 * 7
	expected := 21
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// TestDivision tests the division of two numbers
func TestDivision(t *testing.T) {
	result := 10 / 2
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// TestFailingExample demonstrates a failing test
func TestFailingExample(t *testing.T) {
	result := 1 + 1
	expected := 3 // This is intentionally wrong
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// TestSkippedExample demonstrates a skipped test
func TestSkippedExample(t *testing.T) {
	t.Skip("Skipping this test because it's just an example")
	result := 1 + 1
	expected := 2
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
