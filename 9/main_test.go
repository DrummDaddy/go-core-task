package main

import (
	"math"
	"sync"
	"testing"
)

func TestPipeline(t *testing.T) {
	input := make(chan uint8)
	otput := make(chan float64)

	var wg sync.WaitGroup
	wg.Add(1)

	go StartPipline(input, otput, &wg)

	go func() {
		for _, val := range []uint8{1, 2, 3, 4} {
			input <- val
		}
		close(input)
	}()
	// Ожидаемые результаты
	expected := []float64{1, 8, 27, 64}

	//Чтение из output и проверка результата
	var results []float64
	go func() {
		wg.Wait()
		close(otput)

	}()

	for res := range otput {
		results = append(results, res)
	}

	//Сравниваем результаты
	for i, val := range expected {
		if math.Abs(results[i]-val) > 1e-9 {
			t.Errorf("Ожидалось %.2f, получено %.2f", val, results[i])
		}
	}
}
