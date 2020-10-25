package main

import (
	"fmt"
	"math"
)

func main() {
	var x, epsilon float64
	var n int
	fmt.Println("Разложение eˆx в ряд Тейлора")
	fmt.Printf("Введите x: ")
	fmt.Scan(&x)
	fmt.Printf("Введите степень точности расчета: ")
	fmt.Scan(&n)
	epsilon = 1 / math.Pow10(n)
	var result, prevResult = 0.0, 0.0
	fact := 1
	k := 0
	for {
		fmt.Println("Итерация:")
		fmt.Println("")
		if k > 0 {
			fact *= 2 * k * (2*k + 1)
			fmt.Println("fact:", fact)
		}
		result += math.Pow(x, float64(k)) / float64(fact)
		fmt.Println("result:", result)
		fmt.Println("check:", result-prevResult, "epsilon:", epsilon)
		if math.Abs(result-prevResult) < epsilon {
			fmt.Println("")
			fmt.Println("Итог:", result)
			break
		}
		k++
		prevResult = result
	}
}
