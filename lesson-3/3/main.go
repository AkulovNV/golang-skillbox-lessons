package main

import "fmt"

func main() {

	var a, b int
	fmt.Println("Замена двух целых чисел местами")
	fmt.Printf("Введите число A: ")
	fmt.Scan(&a)
	fmt.Printf("Введите число B: ")
	fmt.Scan(&b)

	fmt.Println("Проивзодим подмену...")
	a, b = b, a
	fmt.Printf("Теперь число A - %d, число B - %d\n", a, b)
}
