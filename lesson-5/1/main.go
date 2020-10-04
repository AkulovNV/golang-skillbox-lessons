package main

import "fmt"

func main() {

	var x, y float64
	var quarter int

	fmt.Println("Определение координатной плоскости по заданным точкам")
	fmt.Printf("Укажите значение x: ")
	fmt.Scan(&x)
	fmt.Printf("Укажите значение y: ")
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		quarter = 1
	} else if x < 0 && y > 0 {
		quarter = 4
	} else if x > 0 && y < 0 {
		quarter = 2
	} else if x < 0 && y < 0 {
		quarter = 3
	} else {
		quarter = 0
	}
	if quarter == 0 {
		fmt.Printf("Заданная точка находится на оси координат\n")
	} else {
		fmt.Printf("Заданная точка находятся в %v коррдинатной плоскоски\n", quarter)
	}

}
