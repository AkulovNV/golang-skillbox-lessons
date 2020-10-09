package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Println("Елочка")
	fmt.Printf("Задайте высоту елочки (пример: 5) ")
	fmt.Scan(&size)

	for size < 0 {
		fmt.Println("Размер не может быть отрицательным, попробуйте еще раз")
		fmt.Printf("Задайте размер (пример: 5) ")
		fmt.Scan(&size)
	}

	for column := 1; column <= size; column++ {
		spaces := size - column
		stars := (size-spaces)*2 - 1
		tree := ""
		for i := 1; i <= spaces; i++ {
			tree += " "
		}
		for i := 1; i <= stars; i++ {
			tree += "*"
		}
		fmt.Printf("%v\n", tree)
	}
}
