package main

import (
	"math/rand"
	"time"
)

// RandomNumberGenerator запускает генератор случайных чисел и отправляет их в канал.
func RandomNumberGenerator(out chan int, stop chan struct{}) {
	rand.Seed(time.Now().UnixNano())

	for {
		select {
		case out <- rand.Intn(100):
		case <-stop:
			close(out)
			return
		}

	}
}

func main() {

	randomNumbers := make(chan int)
	stop := make(chan struct{})

	// Запуск генератора в отдельной горутине
	go RandomNumberGenerator(randomNumbers, stop)

	// Генерируем 10 случайных чисел
	for i := 0; i < 10; i++ {
		num := <-randomNumbers
		println("Получено случайное число:", num)
	}
	// останавливаем генератор
	close(stop)
}
