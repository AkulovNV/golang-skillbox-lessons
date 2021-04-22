/*
Напишите функцию, принимающую на вход массив и возвращающую массив,
в котором элементы идут в обратном порядке по сравнению с исходным
*/

package main

import "fmt"

func reversArr(arr [10]int) (revArr [10]int) {
	for k, v := range arr {
		revArr[len(arr)-1-k] = v
	}
	return
}

func main() {
	var arr [10]int
	var answer string
	for {
		fmt.Printf("Создать свой массив(y/n)? ")
		fmt.Scan(&answer)
		if answer == "y" {
			for i := range arr {
				fmt.Printf("Задайте элемент массива (число): ")
				fmt.Scan(&arr[i])
			}
			break
		} else if answer == "n" {
			arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			break
		} else {
			fmt.Println("Ошибочный ввод, нажмите 'y' или 'n'")
		}
	}
	fmt.Println("Исходный массив:", arr)
	fmt.Println("Реверсивный массив:", reversArr(arr))
}
