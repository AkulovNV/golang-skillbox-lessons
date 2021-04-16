/*
Написать программу, которая на вход получает число, затем с помощью двух функции преобразует его.
Первая умножает, а вторая прибавляет число. Использовать именованные возвращаемые значения
*/

package main

import (
	"fmt"
	"unicode"
)

func multiple(a *float64) float64 {
	*a = *a * *a
	return *a
}

func add(a *float64) float64 {
	*a = *a + *a
	return *a
}

func main() {
	var num float64
	fmt.Printf("Введите число: ")
	fmt.Scan(&num)
	if unicode.IsSymbol(rune(num)) {
		fmt.Println("Введенно не число")

	}
	multiple(&num)
	add(&num)
	fmt.Println("Результат вычислений:", num)
}
