package main

import (
	"reflect"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	tests := []struct {
		a              []int
		b              []int
		expected       bool
		expectedResult []int
	}{
		{
			a:              []int{65, 3, 58, 678, 64},
			b:              []int{64, 2, 3, 43},
			expected:       true,
			expectedResult: []int{64, 3},
		},
		{
			a:              []int{1, 2, 3},
			b:              []int{4, 5, 6},
			expected:       false,
			expectedResult: []int{},
		},
		{
			a:              []int{},
			b:              []int{1, 2, 3},
			expected:       false,
			expectedResult: []int{},
		},
		{
			a:              []int{1, 2, 3},
			b:              []int{},
			expected:       false,
			expectedResult: []int{},
		},
		{
			a:              []int{1, 1, 1, 2, 2},
			b:              []int{2, 2, 3},
			expected:       true,
			expectedResult: []int{2},
		},
	}

	for _, test := range tests {
		resultBool, resultSlice := FindIntersection(test.a, test.b)

		if resultBool != test.expected {
			t.Errorf("Для = %v и b = %v ожидалось %v, но получено %v", test.a, test.b, test.expected, resultBool)
		}
		if !reflect.DeepEqual(resultSlice, test.expectedResult) {
			t.Errorf("Для a = %v и b = %v ожидался результат %v, но получено %v", test.a, test.b, test.expectedResult, resultSlice)
		}
	}
}
