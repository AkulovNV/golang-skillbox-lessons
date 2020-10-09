package main

import (
	"fmt"
)

func main() {
	var day string
	fmt.Println("Рабочие дни недели")
	fmt.Printf("Введите день недели (например вт): ")
	_, _ = fmt.Scan(&day)

	switch day {
	case "пн":
		fmt.Println("Понедельник")
		fallthrough
	case "вт":
		fmt.Println("Вторник")
		fallthrough
	case "ср":
		fmt.Println("Среда")
		fallthrough
	case "чт":
		fmt.Println("Четверг")
		fallthrough
	case "пт":
		fmt.Println("Пятница")
	case "сб":
		fmt.Println("Суббота - выходной")
	case "вс":
		fmt.Println("Воскресенье - выходной")
	default:
		fmt.Println("Допустимые значения: пн, вт, ср,чт, пт, сб, вс")
	}
}
