package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func multi(x int, y int) int {
	return x * y
}

func operator(f1 func(int, int) int, f2 func(int, int) int) {
	n1, n2 := 10, 20
	fmt.Println("Результат умножения функции f2: ", f2(n1, n2))
	fmt.Println("Результат сложения функции f1: ", f1(n1, n2))
}

func main() {
	operator(add, multi)
}
