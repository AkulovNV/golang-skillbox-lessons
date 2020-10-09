package main

import (
	"fmt"
	"math"
)

func main() {
	var counter8, counter16 int
	for i := 1; i < math.MaxUint32; i++ {
		type number1 uint16
		type number2 uint8
		if number1(i) == 0 {
			counter8++
		}
		if number2(i) == 0 {
			counter16++
		}
	}
	fmt.Printf("Количество переполнений для uint8: %d, для uint16: %d\n", counter8, counter16)
}
