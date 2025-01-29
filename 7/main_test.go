package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2

	}()
	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4

	}()
	go func() {
		defer close(ch3)
		ch3 <- 5
		ch3 <- 6
	}()

	merged := merge(ch1, ch2, ch3)

	var result []int
	for val := range merged {
		result = append(result, val)
	}

	expected := []int{1, 2, 3, 4, 5, 6}
	sort.Ints(result)
	sort.Ints(expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

}
