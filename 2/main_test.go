package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}
	result := sliceExample(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("sliceExample failed, expected %v, got %v", expected, result)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	num := 4
	expected := []int{1, 2, 3, 4}
	result := addeElements(input, num)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("addElements failed, expected %v, got %v", expected, result)

	}
}

func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	copy := copySlice(input)
	input[0] = 99
	if reflect.DeepEqual(copy, input) {
		t.Errorf("copySlice failed, copy is affected by original slice modifications")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	index := 2
	expected := []int{1, 2, 4, 5}
	result := removeElement(input, index)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeElement failed, expected %v, got %v", expected, result)
	}

	// Тест с некорректным индексом
	outOfBoundsIndex := 10
	result = removeElement(input, outOfBoundsIndex)
	if !reflect.DeepEqual(result, input) {
		t.Errorf("removeElement failed on out-of-bounds index, expected %v, got %v", input, result)
	}
}
