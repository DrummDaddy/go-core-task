package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "Basic case",
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:     "No elements in slice2",
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},

		{
			name:     "All elements overlap",
			slice1:   []string{"apple", "banana"},
			slice2:   []string{"apple, banana"},
			expected: []string{},
		},
		{
			name:     "Empty slice",
			slice1:   []string{},
			slice2:   []string{"apple", "banana"},
			expected: []string{},
		},
		{
			name:     "Different elements",
			slice1:   []string{"car", "bike", "bus"},
			slice2:   []string{"plane", "ship", "train"},
			expected: []string{"car", "bike", "bus"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Difference(test.slice1, test.slice2)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For slices %v and %v, expected %v but got %v", test.slice1, test.slice2, test.expected, result)
			}
		})
	}
}
