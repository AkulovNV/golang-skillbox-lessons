package main

import (
	"fmt"
)

func main() {
	var month string
	fmt.Println("Времена года")
	fmt.Printf("Введите месяц: ")
	_, _ = fmt.Scan(&month)

	switch month {
	case "декабрь", "январь", "февраль":
		fmt.Println("Зима")
	case "март", "апрель", "май":
		fmt.Println("Весна")
	case "июнь", "июль", "август":
		fmt.Println("Лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("Осень")
	default:
		fmt.Println("Допустимые значения: декабрь, январь, февраль, март, апрель, май, июнь, июль, август, сентябрь, октябрь, ноябрь")
	}
}
