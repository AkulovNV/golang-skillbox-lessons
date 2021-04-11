package main

import (
	"fmt"
	"math"
)

func double(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqrt(y int) float64 {
	return math.Sqrt(float64(y))
}

func operator(f1 func(int) float64, f2 func(int) float64, a, b int) {
	fmt.Println("Результат вычисления функции f1: ", f1(b))
	fmt.Println("Результат вычисления функции f2: ", f2(a))
}

func main() {
	n1, n2 := 10, 20
	operator(double, sqrt, n1, n2)
}
