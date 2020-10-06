package main

import (
	"fmt"
)

func main() {

	var ticket int

	fmt.Println("Счастливый билет")
	fmt.Printf("Введите номер билета: ")
	fmt.Scan(&ticket)

	if ticket/100 == ticket%100 || ticket/100/10+(ticket/100)%10 == (ticket%100/10)+ticket%10 {
		fmt.Println("Поздравляем, у вас счастливый билет :)")
	} else {
		fmt.Println("Не повезло :(")
	}
}
