package main

import "fmt"

func main() {

	var fistNum, secondNum, thirdNum int
	var counter = 0

	fmt.Println("Проверка чисел на значение больше 5")
	fmt.Printf("Введите первое число: ")
	fmt.Scan(&fistNum)
	fmt.Printf("Введите второе число: ")
	fmt.Scan(&secondNum)
	fmt.Printf("Введите третье число: ")
	fmt.Scan(&thirdNum)

	if fistNum > 5 {
		counter++
	}
	if secondNum > 5 {
		counter++
	}
	if thirdNum > 5 {
		counter++
	}
	if counter > 0 {
		fmt.Printf("Среди введенных чисел присутствует %d чисел больше 5\n", counter)
	} else {
		fmt.Println("Среди введенных чисел не ни одного больше 5")
	}
}
