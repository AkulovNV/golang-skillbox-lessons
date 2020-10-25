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
		if k > 1 {
			fact *= k
		}
		result += math.Pow(x, float64(k)) / float64(fact)
		if math.Abs(result-prevResult) < epsilon {
			fmt.Println("")
			fmt.Println("Итог:", result)
			break
		}
		k++
		prevResult = result
	}
}
