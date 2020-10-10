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
	n1 := int32(num1)
	n2 := int32(num2)
	result := n1 * n2
	switch {
	case result < math.MinInt32:
		fmt.Printf("Результат умножения: %d\nПодходящий формат: %T\n", result, int64(result))
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
	case int64(result) <= math.MaxUint32:
		fmt.Printf("Результат умножения: %d\nПодходящий формат: %T\n", result, uint32(result))
	default:
		fmt.Printf("Результат умножения: %d\nПодходящий формат: %T\n", result, uint64(result))
	}
}
