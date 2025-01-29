package main

import "fmt"

//FindIntersection проверяет наличие пересечений и возвращает срез пересекающихся значений
func FindIntersection(a, b []int) (bool, []int) {
	intersection := []int{}
	set := make(map[int]bool)

	//Добавляем элемент первого среза в map
	for _, val := range a {
		set[val] = true
	}
	// Проверяем элементы второго среза
	for _, val := range b {
		if set[val] {
			intersection = append(intersection, val)
			set[val] = false
		}
	}

	hasIntersection := len(intersection) > 0
	return hasIntersection, intersection

}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	hasIntersection, intersection := FindIntersection(a, b)
	fmt.Println("Есть пересечения", hasIntersection)
	fmt.Println("Персечение", intersection)

}
