package main

import "fmt"

// Difference функция для поиска элемнтов в slice 1 которых нет в slice2
func Difference(slice1, slise2 []string) []string {
	// Создаем мапу для проверки уникальности элементов
	set := make(map[string]struct{})
	for _, elem := range slise2 {
		set[elem] = struct{}{}
	}

	// Создаем результат добавления уникальных элементов
	var result []string
	for _, elem := range slice1 {
		if _, exist := set[elem]; !exist {
			result = append(result, elem)
		}

	}
	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	result := Difference(slice1, slice2)
	fmt.Println("Result:", result)
}
