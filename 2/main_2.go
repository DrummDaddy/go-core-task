package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Генерация слайса с 10 случайными числами
func generateSlice() []int {
	rand.Seed(time.Now().UnixNano())
	originalSlice := make([]int, 10)
	for i := 0; i < 10; i++ {
		originalSlice[i] = rand.Intn(100)
	}
	return originalSlice
}

// Функция для фильтрации четных чисел
func sliceExample(slice []int) []int {
	evenSlice := []int{}
	for _, v := range slice {
		if v%2 == 0 {
			evenSlice = append(evenSlice, v)
		}
	}
	return evenSlice
}

// Функция для добавления элементов в конец слайса
func addeElements(slice []int, num int) []int {
	return append(slice, num)
}

// Функция для создания копии слайса
func copySlice(slice []int) []int {
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)
	return copySlice
}

// Функция для удаления элемента по индексу
func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	//Генерация оригинального слайса
	originalSlice := generateSlice()
	fmt.Println("Original Slice:", originalSlice)

	//Используем SliceExample
	evenSlice := sliceExample(originalSlice)
	fmt.Println("Even Slice:", evenSlice)

	// Добавление элемента
	newSlice := addeElements(originalSlice, 42)
	fmt.Println("Slice after adding 42:", newSlice)

	// Копирование слайса
	copiedSlice := copySlice(originalSlice)
	fmt.Println("Copyed Slice:", copiedSlice)

	//Удаление элемента
	indexToRemove := 3 // Удаляем элемент по индексу 3
	removeElementSlice := removeElement(originalSlice, indexToRemove)
	fmt.Println("Slice after removing element at index 3:", removeElementSlice)
}
