package main

import "fmt"

func main() {

	var fistNum, secondNum, thirdNum int

	fmt.Println("Проверка чисел на значение больше 5")
	fmt.Printf("Введите первое число: ")
	fmt.Scan(&fistNum)
	fmt.Printf("Введите второе число: ")
	fmt.Scan(&secondNum)
	fmt.Printf("Введите третье число: ")
	fmt.Scan(&thirdNum)

	if fistNum > 5 || secondNum > 5 || thirdNum > 5 {
		fmt.Println("Среди введенных чисел присутствуют числа больше 5")
	} else {
		fmt.Println("реди введенных чисел не ни одного больше 5")
	}
}
