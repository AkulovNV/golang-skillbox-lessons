package main

import "fmt"

func main() {

	var number int
	fmt.Printf("Введите произвольное целое число: ")
	fmt.Scan(&number)

	for i := 0; i <= number; i++ {
		fmt.Println(i)
	}

}
