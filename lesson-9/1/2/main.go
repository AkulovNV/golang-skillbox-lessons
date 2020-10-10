package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 int16
	fmt.Printf("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Printf("Введите второе число: ")
	fmt.Scan(&num2)
	result := int64(num1) * int64(num2)
	switch {
	case result < math.MinInt16:
		fmt.Printf("Результат умножения: %d\nПодходящий формат: %T\n", result, int32(result))
	case result < math.MinInt8:
		fmt.Printf("Результат умножения: %d\nПодходящий формат: %T\n", result, int16(result))
	case result < 0:
		fmt.Printf("Результат умножения: %d\nПодходящий формат: %T\n", result, int8(result))
	case result >= 0 && result <= math.MaxUint8:
		fmt.Printf("Результат умножения: %d\nПодходящий формат: %T\n", result, uint8(result))
	case result <= math.MaxUint16:
		fmt.Printf("Результат умножения: %d\nПодходящий формат: %T\n", result, uint16(result))
	default:
		fmt.Printf("Результат умножения: %d\nПодходящий формат: %T\n", result, uint32(result))
	}
}
