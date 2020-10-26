/*
Написать программу разбиения числа N на M групп не кратных делителей.

Пояснение: программа запрашивает у пользователя число N,
разбивает все целые числа из диапазона от 1 до N по группам так,
чтобы в одной группе не оказались кратные друг другу числа.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var num, counter int
	// var array, buf []int
	fmt.Printf("Введите число: ")
	fmt.Scan(&num)
	array = getArray(num)

	// for len(array) > 0 {
	// 	// Оставшиеся элементы массива [2,1] разбираем в отдельные группы и увеличивам счетчик
	// 	if len(array) <= 2 {
	// 		for {
	// 			counter++
	// 			fmt.Println("Группа №", counter, ":", array[:1])
	// 			array = getArray(len(array) - 1)
	// 			break
	// 		}
	// 	} else {
	// 		// Вычисляем остаток от деления по модулю между первым элементом массива с остальными
	// 		for k, v := range array[1:] {
	// 			// Если делители не найдены, смещаем массив вперед и запоминаем значение элемента
	// 			if v == 1 {
	// 				buf = append(buf, array[0])
	// 				array = getArray(len(array) - 1)
	// 				break
	// 			}
	// 			// Если делитель найден, увеличиваем счетчик и отображаем элементы массива
	// 			if array[0]%v == 0 {
	// 				counter++
	// 				switch {
	// 				case len(buf) > 0:
	// 					for _, slice := range array[0 : k+1] {
	// 						buf = append(buf, slice)
	// 					}
	// 					fmt.Println("Группа №", counter, ":", buf)
	// 					buf = nil
	// 				default:
	// 					fmt.Println("Группа №", counter, ":", array[0:k+1])
	// 				}
	// 				array = array[k+1:]
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	fmt.Println("Количество групп", counter)

}

// Функция разложения числа на срез элементов
func getArray(n int) []int {
	var a []int
	for i := int(math.Abs(float64(n))); i > 0; i-- {
		a = append(a, i)
	}
	return a
}

// Фукнция подсчета количества массивов
func arrayCounter(n int) int {
	var counter = 0
	var stageNumber = 1.0
	if n == 0 {
		return counter
	} else {
		for i := 1; stageNumber <= math.Abs(float64(n)); i++ {
			stageNumber = math.Pow(2.0, float64(i))
			fmt.Println(stageNumber)
			counter++
		}
	}
	return counter
}
