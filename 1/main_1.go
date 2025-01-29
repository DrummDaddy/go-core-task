package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

//Определение переменных различных типов

func defineVariables() (int, int, int, float64, string, bool, complex64) {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	return numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum

}

// Определение типов переменных
func getVariableTypes(vars ...interface{}) string {
	var types []string
	for _, v := range vars {
		types = append(types, reflect.TypeOf(v).String())
	}
	return strings.Join(types, ", ")
}

// Преобразовываем переменные в строку
func converToStrings(vars ...interface{}) string {
	var result []string
	for _, v := range vars {
		result = append(result, fmt.Sprintf("%v", v))
	}
	return strings.Join(result, " ")

}

// Преобразование строки в срез рун
func converToRunes(s string) []rune {
	return []rune(s)
}

// Хэширование среза рун с добавлением соли
func hashWithSalt(runes []rune, salt string) string {
	middle := len(runes) / 2
	saltedRunes := append(append(runes[:middle], []rune(salt)...), runes[middle:]...)
	hasher := sha256.New()
	hasher.Write([]byte(string(saltedRunes)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	// 1. Создаем переменные
	numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum := defineVariables()

	// 2. Определяем типы
	variableTypes := getVariableTypes(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("Типы переменных: ", variableTypes)

	//3. Преобразуем переменные в строку
	combinedString := converToStrings(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("Объединенная строка: ", combinedString)

	//4. Преобразуем строку в срез рун
	runes := converToRunes(combinedString)
	fmt.Println("Срез рун: ", runes)

	//5. Добавляем соль и хеширование
	hash := hashWithSalt(runes, "go-2024")
	fmt.Println("SHA256 хэш с солью 'go-2024': ", hash)

}
