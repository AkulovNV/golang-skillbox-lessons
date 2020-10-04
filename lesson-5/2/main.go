package main

import "fmt"

func main() {

	var fistNum, secondNum, thirdNum int

	fmt.Println("Проверка ввода положительного числа")
	fmt.Printf("Введите первое число: ")
	fmt.Scan(&fistNum)
	fmt.Printf("Введите второе число: ")
	fmt.Scan(&secondNum)
	fmt.Printf("Введите третье число: ")
	fmt.Scan(&thirdNum)

	if fistNum > 0 || secondNum > 0 || thirdNum > 0 {
		fmt.Println("Среди заданных чисел имеются положительные")
	} else {
		fmt.Println("Среди заданных чисел отсутствуют положительные")
	}
}
