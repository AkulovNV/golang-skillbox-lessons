/*
Написать программу, в которой будет 3 функции, попарно использующие разные глобальные переменные.
Функции, должны прибавлять к поданном на вход числу глобальную переменную и возвращать результат.
Затем вызвать по очереди три функции, передавая результат из одной в другую.
*/

package main

import "fmt"

const num1 = 100
const num2 = 55
const num3 = 235

func sum1(a *int) int {
	*a += num1
	return *a
}
func sum2(a *int) int {
	*a += num2
	return *a
}
func sum3(a *int) int {
	*a += num3
	return *a
}
func main() {
	var n int
	fmt.Printf("Введите число: ")
	fmt.Scan(&n)
	sum1(&n)
	sum2(&n)
	sum3(&n)
	fmt.Printf("Глобальные переменные: n1=%d, n2=%d, n3=%d\n", num1, num2, num3)
	fmt.Println("Результат вычислений:", n)
}
