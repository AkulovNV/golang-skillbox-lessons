package main

import "fmt"

func main() {

	var firstNum, secondNum int
	fmt.Printf("Введите первое число: ")
	fmt.Scan(&firstNum)
	fmt.Printf("Введите второе число: ")
	fmt.Scan(&secondNum)

	for i := firstNum; i <= firstNum+secondNum; i++ {
		fmt.Println(i)
	}

}
