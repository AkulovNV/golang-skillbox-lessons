/*
Напишите функцию, принимающую на вход массив и возвращающую массив,
в котором элементы идут в обратном порядке по сравнению с исходным
*/

package main

import "fmt"

func main() {
	var arr [10]int
	var reversArr [10]int
	for i := range arr {
		fmt.Printf("Введите целое число: ")
		fmt.Scan(&arr[i])
		reversArr[len(arr)-1-i] = arr[i]
	}
	fmt.Println("Исходный массив:", arr)
	fmt.Println("Реверсивный массив:", reversArr)
}
