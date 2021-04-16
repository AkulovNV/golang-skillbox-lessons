/*
Написать программу, в которой будет 3 функции, попарно использующие разные глобальные переменные.
Функции, должны прибавлять к поданном на вход числу глобальную переменную и возвращать результат.
Затем вызвать по очереди три функции, передавая результат из одной в другую.
*/

package main

import "fmt"

var (
	num1 = 100
	num2 = 55
	num3 = 235
)

func sum1(a int) (sum int) {
	sum = a + num1 + num2
	return
}
func sum2(a int) (sum int) {
	sum = a + num2 + num3
	return
}
func sum3(a int) (sum int) {
	sum = a + num3 + num1
	return
}
func main() {
	var n int
	fmt.Printf("Введите число: ")
	fmt.Scan(&n)
	fmt.Printf("Глобальные переменные: n1=%d, n2=%d, n3=%d\n", num1, num2, num3)
	fmt.Println("Результат вычислений:", sum1(sum2(sum3(n))))
}
