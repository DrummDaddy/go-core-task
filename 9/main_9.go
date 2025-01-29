package main

import (
	"fmt"
	"math"
	"sync"
)

// Конвейер: Чтенгие чисел из первого канала, преобразование и запись во второй канал
func StartPipline(input <-chan uint8, output chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range input {
		//Преобразование во float и возведение в куб
		cubed := math.Pow(float64(num), 3)
		output <- cubed
	}
}

func main() {
	// Создаем каналы
	input := make(chan uint8)
	output := make(chan float64)

	var wg sync.WaitGroup

	// Запускаем конвейер в отдельной горутине
	wg.Add(1)
	go StartPipline(input, output, &wg)
	// Запускаем данные в первый канал
	go func() {
		for i := uint8(1); i <= 10; i++ {
			input <- i
		}
		close(input)
	}()
	//Читаем данные из второго канала в глафной функции
	go func() {
		wg.Wait()
		close(output)

	}()
	//Вывод результатов
	fmt.Println("Результаты:")
	for result := range output {
		fmt.Println(result)
	}
}
