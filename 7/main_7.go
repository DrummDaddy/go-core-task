package main

import (
	"fmt"
	"sync"
)

// Функция для объединения нескольких каналов в один
func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	//Анонимная функция для чтения данных из канала
	multiplexer := func(ch <-chan int) {
		defer wg.Done()
		for val := range ch {
			out <- val
		}
	}
	// Добавляем количество каналов в группу ожидания
	wg.Add(len(channels))

	for _, ch := range channels {
		go multiplexer(ch)
	}

	//Закрываем резльтирующий канал, когда все горутины завершатся
	go func() {
		wg.Wait()
		close(out)

	}()
	return out

}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch2 <- 2
		ch3 <- 3

	}()
	go func() {
		defer close(ch2)
		ch2 <- 4
		ch2 <- 5

	}()
	go func() {
		defer close(ch3)
		ch3 <- 6
		ch3 <- 7
		ch3 <- 8

	}()

	merged := merge(ch1, ch2, ch3)

	for val := range merged {
		fmt.Println(val)
	}
}
