package main

import "fmt"

func main() {

	var fistNum, secondNum, thirdNum int

	fmt.Println("Проверка одинаковых чисел")
	fmt.Printf("Введите первое число: ")
	fmt.Scan(&fistNum)
	fmt.Printf("Введите второе число: ")
	fmt.Scan(&secondNum)
	fmt.Printf("Введите третье число: ")
	fmt.Scan(&thirdNum)

	if fistNum == secondNum || secondNum == thirdNum || thirdNum == fistNum {
		fmt.Println("Среди заданных чисел имеются одинаковые")
	} else {
		fmt.Println("Среди заданных чисел отсутствуют одиковые")
	}
}
