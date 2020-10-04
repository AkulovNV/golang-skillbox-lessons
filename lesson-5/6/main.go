package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c float64

	fmt.Println("Вычисление квадратного уравнения: axˆ2 + bx + c = 0 ")
	fmt.Printf("Введите коэффициент a: ")
	fmt.Scan(&a)
	fmt.Printf("Введите коэффициент b: ")
	fmt.Scan(&b)
	fmt.Printf("Введите коэффициент c: ")
	fmt.Scan(&c)

	if a == 0 {
		fmt.Println("коэффициент a не может быть нулевым")
	} else {
		D := math.Pow(b, 2) - 4*a*c
		if D > 0 {
			x1 := (-b + math.Sqrt(D)) / (2 * a)
			x2 := (-b - math.Sqrt(D)) / (2 * a)
			fmt.Printf("Дискриминант больше нуля: x1 = %.1f, x2 = %.1f\n", x1, x2)
		} else if D == 0 {
			x := (-b + math.Sqrt(D)) / (2 * a)
			fmt.Printf("Дискриминант равен нулю: x = %.1f\n", x)
		} else {
			fmt.Println("Дискриминант меньше нуля: корней нет")
		}
	}
}
