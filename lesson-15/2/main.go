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
	for i := range arr {
		fmt.Printf("Задайте элемент массива (число): ")
		fmt.Scan(&arr[i])
	}
	fmt.Println("Исходный массив:", arr)
	fmt.Println("Реверсивный массив:", reversArr(arr))
}
