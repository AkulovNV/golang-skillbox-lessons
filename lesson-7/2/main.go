package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Println("Шахматная доска")
	fmt.Printf("Задайте размер (пример: 8) ")
	fmt.Scan(&size)

	white := " "
	black := "*"

	for size < 0 {
		fmt.Println("Размер не может быть отрицательным, попробуйте еще раз")
		fmt.Printf("Задайте размер (пример: 8) ")
		fmt.Scan(&size)
	}

	for column := 1; column <= size; column++ {
		row := ""
		if column%2 > 0 {
			row = " "
		}
		for i := 1; i <= size; i++ {
			if i%2 == 0 {
				row += black
			} else {
				row += white
			}
		}
		fmt.Printf("%v\n", row)
	}
}
