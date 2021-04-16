/*
Написать программу, которая на вход получается число.
Далее число подается в функцию, которая возвращает true, если число четное, и false, если нечетное.
Вывести в консоль результат функции.
*/

package main

import (
	"fmt"
)

func checkParity(a int) bool {
	return *a%2 == 0
}

func main() {
	var num int
	fmt.Printf("Введите целое число: ")
	fmt.Scan(&num)
	if checkParity(num) == true {
		fmt.Println("четное")
	} else {
		fmt.Println("нечетное")
	}
}
