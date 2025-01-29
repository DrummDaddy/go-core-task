package main

import (
	"testing"
	"time"
)

func TestRandomNumberGenerator(t *testing.T) {
	randomNumbers := make(chan int)

	stop := make(chan struct{})

	go RandomNumberGenerator(randomNumbers, stop)

	// считываем первые 5 чисел
	numbers := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		select {
		case num := <-randomNumbers:
			numbers = append(numbers, num)
		case <-time.After(time.Second): // на случай если завис генератор
			t.Fatal("Тест не дождался случайного числа")
		}
	}

	if len(numbers) != 5 {
		t.Errorf("Ожидаемое количество чисел: 5, полчено: %d", len(numbers))
	}

	for _, num := range numbers {
		if num < 0 || num >= 100 {
			t.Errorf("Число %d вне диапазона [0, 99]", num)
		}
	}

	close(stop)
}

func TestStopGenerator(t *testing.T) {
	randomNumbers := make(chan int)
	stop := make(chan struct{})

	go RandomNumberGenerator(randomNumbers, stop)

	close(stop)

	_, ok := <-randomNumbers
	if ok {
		t.Error("Канал не был закрыт после остановки генератора")
	}
}
