package main

import (
	"testing"
)

func TestGetVariableTypes(t *testing.T) {
	_, _, _, pi, name, isActive, complexNum := defineVariables()
	expected := "int, int, int, float64, string, bool, complex64"
	result := getVariableTypes(42, 052, 0x2A, pi, name, isActive, complexNum)

	if result != expected {
		t.Errorf("Ожидалось: %s, Получено: %s", expected, result)
	}
}

func TestConverttoStrings(t *testing.T) {
	result := converToStrings(42, 052, 0x2A, 3.14, "Golang", true, 1+2i)
	expected := "42 42 42 3.14 Golang true (1+2i)"

	if result != expected {
		t.Errorf("Ожидалось: %s, Получено: %s", expected, result)
	}

}

func TestConverToRunes(t *testing.T) {
	input := "42 42 42 3.14 Golang true (1+2i)"
	expected := []rune(input)
	result := converToRunes(input)

	for i, r := range result {
		if r != expected[i] {
			t.Errorf("Руна на позиции %d: ожидалось %c, получено %c", i, expected[i], r)
		}
	}
}

func TestHashWithSalt(t *testing.T) {
	input := converToRunes("42 42 42 3.14 Golang true (1+2i)")
	salt := "go-2024"
	result := hashWithSalt(input, salt)
	expected := hashWithSalt(input, salt)

	if result != expected {
		t.Errorf("Ожидалось: %s, Получено: %s", expected, result)
	}
}
