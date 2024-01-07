package testmymath

import (
	"testing"
)

func TestTestMySqrt(t *testing.T) {
	x := 4.0
	expected := 2.0
	result := TestMySqrt(x)
	if result != expected {
		t.Errorf("MySqrt(%v) = %v; expected %v", x, result, expected)
	}
}

func TestTestMyCeil(t *testing.T) {
	x := 3.5
	expected := 4.0
	result := TestMyCeil(x)

	if result != expected {
		t.Errorf("MyCeil(%v) = %v; expected %v", x, result, expected)
	}
}

func TestTestMyFloor(t *testing.T) {
	x := 3.5
	expected := 3.0
	result := TestMyFloor(x)
	if result != expected {
		t.Errorf("MyFloor(%v) = %v; expected %v", x, result, expected)
	}
}

func TestTestMyPow(t *testing.T) {
	x := 2.0
	y := 3.0
	expected := 8.0
	result := TestMyPow(x, y)
	if result != expected {
		t.Errorf("MyPow(%v, %v) = %v; expected %v", x, y, result, expected)
	}
}

func TestTestMyMax(t *testing.T) {
	x := 2.0
	y := 3.0
	expected := 3.0
	result := TestMyMax(x, y)

	if result != expected {
		t.Errorf("MyMax(%v, %v) = %v; expected %v", x, y, result, expected)
	}
}

func TestTestMyMin(t *testing.T) {
	x := 2.0
	y := 3.0
	expected := 2.0
	result := TestMyMin(x, y)

	if result != expected {
		t.Errorf("MyMin(%v, %v) = %v; expected %v", x, y, result, expected)
	}
}
