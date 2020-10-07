// Данный вариант оптимизирует работу лифта за счет сокращения движения до нужного этажа, вместо максимального

package main

import "fmt"

func main() {

	fmt.Println("Движение лифта")
	var maxFloor int
	var stop = false
	var minFloor = 1
	var currentFloor = 1

	var direction = 1
	var maxPassengers = 2
	var totalPassengers = 0

	var passenger1 = 4
	var passenger2 = 7
	var passenger3 = 10

	for {

		// Определяем максимальный этаж, с которого вызвали лифт
		if passenger1 > passenger2 && passenger1 > passenger3 {
			maxFloor = passenger1
		} else if passenger2 > passenger1 && passenger2 > passenger3 {
			maxFloor = passenger2
		} else {
			maxFloor = passenger3
		}

		// Определяем направление движения лифта
		if currentFloor == maxFloor {
			direction = -1
		} else if currentFloor == minFloor {
			direction = 1
			totalPassengers = 0
		}

		// Определяем должен ли остановиться лифт, чтобы забрать пассажиров
		if totalPassengers < maxPassengers && direction == -1 {
			stop = true
		} else {
			stop = false
		}

		// Определяем есть ли пассажиры на этаже и можем ли мы их забрать
		if currentFloor == passenger1 && stop {
			fmt.Printf("Забираем пассажира с этажа %v\n", currentFloor)
			passenger1 = 1
			totalPassengers++
		}
		if currentFloor == passenger2 && stop {
			fmt.Printf("Забираем пассажира с этажа %v\n", currentFloor)
			passenger2 = 1
			totalPassengers++
		}
		if currentFloor == passenger3 && stop {
			fmt.Printf("Забираем пассажира с этажа %v\n", currentFloor)
			passenger3 = 1
			totalPassengers++
		}

		currentFloor += direction
		fmt.Printf("Лифт находится на этаже: %v, Количество пассажиров: %v\n", currentFloor, totalPassengers)

		// Если все пассажиры доставлены, завершам работу
		if passenger1 == 1 && passenger2 == 1 && passenger3 == 1 && currentFloor == minFloor {
			break
		}
	}
}
