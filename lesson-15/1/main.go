/*
Разработайте программу, позволяющую ввести 10 целых чисел,
а затем вывести из них количество чётных и нечётных чисел.
Для ввода и подсчёта используйте разные циклы.
*/

package main

import "fmt"

func main() {
	var numbers [10]int
	var even, odd int

	for i := range numbers {
		fmt.Printf("Введите целое число: ")
		fmt.Scan(&numbers[i])
	}
	fmt.Println("================")
	for _, n := range numbers {
		if n%2 == 0 {
			even += 1
			fmt.Printf("Число %d является четным\n", n)
		} else {
			odd += 1
			fmt.Printf("Число %d является нечетным\n", n)
		}
	}
	fmt.Println("Количество четных чисел в массиве:", even)
	fmt.Println("Количество не четных чисел в массиве:", odd)
}
