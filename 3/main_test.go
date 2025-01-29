package main

import (
	"testing"
)

func TestStringIntMap(t *testing.T) {
	m := NewStringIntMap()
	// тестируем метод Add
	m.Add("a", 10)
	m.Add("b", 20)
	m.Add("c", 30)
	if len(m.items) != 3 {
		t.Errorf("Expected 3 elements, got %d elements", len(m.items))
	}
	// тест метода Get
	value, exists := m.Get("a")
	if !exists || value != 10 {
		t.Errorf("Expected value 10 for key a, got %d", value)
	}

	// тест метода Remove
	m.Remove("b")
	_, exists = m.Get("b")
	if exists {
		t.Errorf("Expected key 'b' to be deleted")
	}
	if len(m.items) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(m.items))
	}
	// тест метода Exists
	if !m.Exists("a") {
		t.Errorf("Expected key a to exist")

	}

	if m.Exists("b") {
		t.Errorf("Expected key b are not exists")
	}

	// тест метода Copy
	copyMap := m.Copy()
	if len(copyMap) != 2 {
		t.Errorf("Expected copy map to have 2 elements, got %d", len(copyMap))
	}
	if copyMap["a"] != 10 || copyMap["c"] != 30 {
		t.Errorf("Values in copy map do not match")
	}

}
